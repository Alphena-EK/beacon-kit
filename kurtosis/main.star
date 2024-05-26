el_cl_genesis_data_generator = import_module(
    "github.com/kurtosis-tech/ethereum-package/src/prelaunch_data_generator/el_cl_genesis/el_cl_genesis_generator.star",
)

execution = import_module("./src/nodes/execution/execution.star")

beacond = import_module("./src/nodes/consensus/beacond/launcher.star")
networks = import_module("./src/networks/networks.star")
port_spec_lib = import_module("./src/lib/port_spec.star")
nodes = import_module("./src/nodes/nodes.star")
nginx = import_module("./src/services/nginx/nginx.star")
constants = import_module("./src/constants.star")
goomy_blob = import_module("./src/services/goomy/launcher.star")
prometheus = import_module("./src/observability/prometheus/prometheus.star")
grafana = import_module("./src/observability/grafana/grafana.star")
pyroscope = import_module("./src/observability/pyroscope/pyroscope.star")
tx_fuzz = import_module("./src/services/tx_fuzz/launcher.star")

def run(plan, validators, full_nodes = [], seed_nodes = [], rpc_endpoints = [], boot_sequence = {"type": "sequential"}, additional_services = [], metrics_enabled_services = []):
    """
    Initiates the execution plan with the specified number of validators and arguments.

    Args:
    plan: The execution plan to be run.
    args: Additional arguments to configure the plan. Defaults to an empty dictionary.
    """

    validators = nodes.parse_nodes_from_dict(validators)
    full_nodes = nodes.parse_nodes_from_dict(full_nodes)
    seed_nodes = nodes.parse_nodes_from_dict(seed_nodes)
    num_validators = len(validators)
    bootnode_count = num_validators * 0.15
    bootnode_count = int(bootnode_count)
    if bootnode_count < 1:
        bootnode_count = 1

    # 1. Initialize EVM genesis data
    evm_genesis_data = networks.get_genesis_data(plan)

    node_modules = {}
    for node in validators:
        if node.el_type not in node_modules.keys():
            node_path = "./src/nodes/execution/{}/config.star".format(node.el_type)
            node_module = import_module(node_path)
            node_modules[node.el_type] = node_module

    for node in full_nodes:
        if node.el_type not in node_modules.keys():
            node_path = "./src/nodes/execution/{}/config.star".format(node.el_type)
            node_module = import_module(node_path)
            node_modules[node.el_type] = node_module

    for node in seed_nodes:
        if node.el_type not in node_modules.keys():
            node_path = "./src/nodes/execution/{}/config.star".format(node.el_type)
            node_module = import_module(node_path)
            node_modules[node.el_type] = node_module

    # 2. Upload files
    jwt_file, kzg_trusted_setup = execution.upload_global_files(plan, node_modules)

    # 3. Perform genesis ceremony
    if boot_sequence["type"] == "sequential":
        beacond.perform_genesis_ceremony(plan, validators, jwt_file)
    else:
        beacond.perform_genesis_ceremony_parallel(plan, validators, jwt_file)

    el_enode_addrs = []
    metrics_enabled_services = metrics_enabled_services[:]

    consensus_node_peering_info = []
    all_consensus_peering_info = {}

    # 4i. Start seed nodes
    seed_node_el_client_configs = []
    for n, seed in enumerate(seed_nodes):
        el_client_config = execution.generate_node_config(plan, node_modules, seed, "seed", n)
        seed_node_el_client_configs.append(el_client_config)

    if seed_node_el_client_configs != []:
        seed_node_el_clients = execution.deploy_nodes(plan, seed_node_el_client_configs)

    for n, seed in enumerate(seed_nodes):
        el_service_name = "el-{}-{}-{}".format("seed", seed.el_type, n)
        enode_addr = execution.get_enode_addr(plan, el_service_name)
        el_enode_addrs.append(enode_addr)
        if seed.el_type != "ethereumjs":
            metrics_enabled_services.append({
                "name": el_service_name,
                "service": seed_node_el_clients[el_service_name],
                "metrics_path": node_modules[seed.el_type].METRICS_PATH,
            })

    seed_node_configs = {}
    for n, seed in enumerate(seed_nodes):
        cl_service_name = "cl-seed-beaconkit-{}".format(n)
        el_client = "el-{}-{}-{}".format("seed", seed.el_type, n)
        seed_node_config = beacond.create_node_config(plan, seed.cl_image, consensus_node_peering_info, el_client, "seed", jwt_file, kzg_trusted_setup, n)
        seed_node_configs[cl_service_name] = seed_node_config

    seed_nodes_clients = plan.add_services(
            configs = seed_node_configs,
        )

    for n, seed_client in enumerate(seed_nodes):
            cl_service_name = "cl-seed-beaconkit-{}".format(n)
            peer_info = beacond.get_peer_info(plan, cl_service_name)
            consensus_node_peering_info.append(peer_info)

            if seed_client.el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": cl_service_name,
                    "service": seed_nodes_clients[cl_service_name],
                    "metrics_path": beacond.METRICS_PATH,
                })

    # 4. Start network validators
    validator_node_el_clients = []
    if boot_sequence["type"] == "sequential":
        for n, validator in enumerate(validators):
            el_client = execution.create_node(plan, node_modules, validator, "validator", n, el_enode_addrs)
            validator_node_el_clients.append(el_client["service"])
            el_enode_addrs.append(el_client["enode_addr"])

            # As ethereumjs currently does not support metrics, we only add the metrics path for other clients
            if validator.el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": el_client["name"],
                    "service": el_client["service"],
                    "metrics_path": node_modules[validator.el_type].METRICS_PATH,
                })

            # 4b. Launch CL
            beacond_service = beacond.create_node(plan, validator.cl_image, consensus_node_peering_info[:n], el_client["name"], jwt_file, kzg_trusted_setup, n)
            peer_info = beacond.get_peer_info(plan, beacond_service.name)
            consensus_node_peering_info.append(peer_info)
            if validator.el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": beacond_service.name,
                    "service": beacond_service,
                    "metrics_path": beacond.METRICS_PATH,
                })
    else:
        bootnode_validators = validators[:bootnode_count]
        remaining_validators = validators[bootnode_count:]
        el_client_configs = []
        for n, validator in enumerate(bootnode_validators):
            el_client_config = execution.generate_node_config(plan, node_modules, validator, "validator", n, el_enode_addrs)
            el_client_configs.append(el_client_config)
        
        el_clients = execution.deploy_nodes(plan, el_client_configs)

        for n, el_client in enumerate(bootnode_validators):
            el_service_name = "el-{}-{}-{}".format("validator", el_client.el_type, n)
            validator_node_el_clients.append(el_service_name)
            if validators[n].el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": el_service_name,
                    "service": el_clients[el_service_name],
                    "metrics_path": node_modules[bootnode_validators[n].el_type].METRICS_PATH,
                })

        validator_node_configs = {}
        for n, validator in enumerate(bootnode_validators):
            cl_service_name = "cl-validator-beaconkit-{}".format(n)
            el_client = el_clients.keys()[n]
            validator_node_config = beacond.create_node_config(plan, validator.cl_image, consensus_node_peering_info, el_client, "validator", jwt_file, kzg_trusted_setup, n)
            validator_node_configs[cl_service_name] = validator_node_config
        
        cl_clients = plan.add_services(
            configs = validator_node_configs,
        )

        for n, cl_client in enumerate(bootnode_validators):
            cl_service_name = "cl-validator-beaconkit-{}".format(n)
            peer_info = beacond.get_peer_info(plan, cl_service_name)
            all_consensus_peering_info[cl_service_name] = peer_info
            # consensus_node_peering_info.append(peer_info)

            if validators[n].el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": cl_service_name,
                    "service": cl_clients[cl_service_name],
                    "metrics_path": beacond.METRICS_PATH,
                })
        
        remaining_el_client_configs = []
        for n, validator in enumerate(remaining_validators):
            index = n+bootnode_count
            el_client_config = execution.generate_node_config(plan, node_modules, validator, "validator", index, el_enode_addrs)
            remaining_el_client_configs.append(el_client_config)

        remaining_el_clients = execution.deploy_nodes(plan, remaining_el_client_configs)

        for n, validator in enumerate(remaining_validators):
            index = n+bootnode_count
            el_service_name = "el-{}-{}-{}".format("validator", validator.el_type, index)
            if validator.el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": el_service_name,
                    "service": remaining_el_clients[el_service_name],
                    "metrics_path": node_modules[validator.el_type].METRICS_PATH,
                })

        validator_node_configs = {}
        for n, validator in enumerate(remaining_validators):
            index = n+bootnode_count
            cl_service_name = "cl-validator-beaconkit-{}".format(index)
            el_client = "el-{}-{}-{}".format("validator", validator.el_type, index)
            validator_node_config = beacond.create_node_config(plan, validator.cl_image, consensus_node_peering_info, el_client, "validator", jwt_file, kzg_trusted_setup, index)
            validator_node_configs[cl_service_name] = validator_node_config
        
        remaining_cl_clients = plan.add_services(
            configs = validator_node_configs,
        )

        for n, validator in enumerate(remaining_validators):
            index = n+bootnode_count
            cl_service_name = "cl-validator-beaconkit-{}".format(index)
            peer_info = beacond.get_peer_info(plan, cl_service_name)
            all_consensus_peering_info[cl_service_name] = peer_info
            if validator.el_type != "ethereumjs":
                metrics_enabled_services.append({
                    "name": cl_service_name,
                    "service": remaining_cl_clients[cl_service_name],
                    "metrics_path": beacond.METRICS_PATH,
                })

    for n, seed_node in enumerate(seed_nodes):
        cl_service_name = "cl-seed-beaconkit-{}".format(n)
        beacond.dial_unsafe_peers(plan, cl_service_name, all_consensus_peering_info)

    # 5. Start full nodes (rpcs)
    full_node_configs = {}
    full_node_el_client_configs = []
    for n, full in enumerate(full_nodes):
        el_client_config = execution.generate_node_config(plan, node_modules, full, "full", n, el_enode_addrs)
        full_node_el_client_configs.append(el_client_config)

    if full_node_el_client_configs != []:
        full_node_el_clients = execution.deploy_nodes(plan, full_node_el_client_configs)

    for n, full in enumerate(full_nodes):
        el_client = "el-{}-{}-{}".format("full", full.el_type, n)
        if full.el_type != "ethereumjs":
            metrics_enabled_services.append({
                "name": el_client,
                "service": full_node_el_clients[el_client],
                "metrics_path": node_modules[full.el_type].METRICS_PATH,
            })

    for n, full in enumerate(full_nodes):
        # 5b. Launch CL
        cl_service_name = "cl-full-beaconkit-{}".format(n)
        el_client = "el-{}-{}-{}".format("full", full.el_type, n)
        full_node_config = beacond.create_node_config(plan, full.cl_image, consensus_node_peering_info, el_client, "full", jwt_file, kzg_trusted_setup, n)
        full_node_configs[cl_service_name] = full_node_config

    if full_node_configs != {}:
        services = plan.add_services(
            configs = full_node_configs,
        )

    for name, service in services.items():
        # excluding ethereumjs from metrics as it is the last full node in the args file beaconkit-all.yaml, TO-DO: to improve this later
        if name != cl_service_name:
            metrics_enabled_services.append({
                "name": name,
                "service": service,
                "metrics_path": beacond.METRICS_PATH,
            })

    # 6. Start RPCs
    for n, rpc in enumerate(rpc_endpoints):
        nginx.get_config(plan, rpc["services"])

    # 7. Start additional services
    for s in additional_services:
        if s == "goomy_blob":
            plan.print("Launching Goomy the Blob Spammer")
            goomy_blob_args = {"goomy_blob_args": []}
            goomy_blob.launch_goomy_blob(
                plan,
                constants.PRE_FUNDED_ACCOUNTS[0],
                plan.get_service("nginx").ports["http"].url,
                goomy_blob_args,
            )
            plan.print("Successfully launched goomy the blob spammer")

    if "tx-fuzz" in additional_services:
        plan.print("Launching tx-fuzz")
        fuzzing_node = validator_node_el_clients[0]
        if len(full_nodes) > 0:
            fuzzing_node = full_node_el_clients[full_node_el_clients.keys()[0]]
        tx_fuzz.launch_tx_fuzz(
            plan,
            constants.PRE_FUNDED_ACCOUNTS[1].private_key,
            "http://{}:{}".format(fuzzing_node.ip_address, execution.RPC_PORT_NUM),
            [],
        )

    if "prometheus" in additional_services:
        prometheus_url = prometheus.start(plan, metrics_enabled_services)

        if "grafana" in additional_services:
            grafana.start(plan, prometheus_url)

        if "pyroscope" in additional_services:
            pyroscope.run(plan)

    plan.print("Successfully launched development network")
