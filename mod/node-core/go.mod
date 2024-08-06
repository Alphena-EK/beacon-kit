module github.com/berachain/beacon-kit/mod/node-core

go 1.22.5

replace (
	// The following are required to build with the latest version of the cosmos-sdk main branch:
	cosmossdk.io/api => cosmossdk.io/api v0.7.3-0.20240731205446-aee9803a0af6
	cosmossdk.io/core => cosmossdk.io/core v0.12.1-0.20240731205446-aee9803a0af6
	cosmossdk.io/core/testing => cosmossdk.io/core/testing v0.0.0-20240731202123-43dd23137e9d
	github.com/berachain/beacon-kit/mod/config => ../config
	github.com/berachain/beacon-kit/mod/consensus => ../consensus
	github.com/berachain/beacon-kit/mod/consensus-types => ../consensus-types
	github.com/berachain/beacon-kit/mod/depinject => ../depinject
	github.com/berachain/beacon-kit/mod/execution => ../execution
	github.com/berachain/beacon-kit/mod/log => ../log
	github.com/berachain/beacon-kit/mod/runtime => ../runtime
	github.com/berachain/beacon-kit/mod/state-transition => ../state-transition
	github.com/berachain/beacon-kit/mod/storage => ../storage
	github.com/cosmos/cosmos-sdk => github.com/berachain/cosmos-sdk v0.46.0-beta2.0.20240729221345-3df18b5c3b34
)

require (
	cosmossdk.io/core v0.12.1-0.20240726151427-2e0884564fdb
	cosmossdk.io/log v1.3.2-0.20240530141513-465410c75bce
	cosmossdk.io/store/v2 v2.0.0-20240729225221-c4de9a970633
	github.com/berachain/beacon-kit/mod/async v0.0.0-20240705193247-d464364483df
	github.com/berachain/beacon-kit/mod/beacon v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/config v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/consensus v0.0.0-20240723155519-565f208d5482
	github.com/berachain/beacon-kit/mod/consensus-types v0.0.0-20240806160829-cde2d1347e7e
	github.com/berachain/beacon-kit/mod/da v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/depinject v0.0.0-00010101000000-000000000000
	github.com/berachain/beacon-kit/mod/engine-primitives v0.0.0-20240806160829-cde2d1347e7e
	github.com/berachain/beacon-kit/mod/errors v0.0.0-20240731201523-f427600b4713
	github.com/berachain/beacon-kit/mod/execution v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/geth-primitives v0.0.0-20240806160829-cde2d1347e7e
	github.com/berachain/beacon-kit/mod/log v0.0.0-20240726221339-a8bfeebf8ecf
	github.com/berachain/beacon-kit/mod/node-api v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/node-api/engines v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/payload v0.0.0-20240728004405-58f868b44614
	github.com/berachain/beacon-kit/mod/primitives v0.0.0-20240806160829-cde2d1347e7e
	github.com/berachain/beacon-kit/mod/runtime v0.0.0-20240731201523-f427600b4713
	github.com/berachain/beacon-kit/mod/state-transition v0.0.0-20240717225334-64ec6650da31
	github.com/berachain/beacon-kit/mod/storage v0.0.0-20240728004405-58f868b44614
	github.com/cometbft/cometbft v1.0.0-rc1.0.20240729121641-d06d2e8229ee
	github.com/cosmos/cosmos-sdk v0.53.0
	github.com/crate-crypto/go-kzg-4844 v1.1.0
	github.com/hashicorp/go-metrics v0.5.3
	github.com/itsdevbear/comet-bls12-381 v0.0.0-20240413212931-2ae2f204cde7
	github.com/spf13/afero v1.11.0
)

require (
	cosmossdk.io/collections v0.4.0 // indirect
	cosmossdk.io/core/testing v0.0.0-00010101000000-000000000000 // indirect
	cosmossdk.io/errors v1.0.1 // indirect
	cosmossdk.io/server/v2/stf v0.0.0-20240802204200-bc0731cb49e9 // indirect
	cosmossdk.io/x/auth v0.0.0-20240731202123-43dd23137e9d // indirect
	github.com/DataDog/datadog-go v4.8.3+incompatible // indirect
	github.com/DataDog/zstd v1.5.6 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/VictoriaMetrics/fastcache v1.12.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/berachain/beacon-kit/mod/chain-spec v0.0.0-20240705193247-d464364483df // indirect
	github.com/berachain/beacon-kit/mod/p2p v0.0.0-20240618214413-d5ec0e66b3dd // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240616162244-4768e80dfb9a // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v1.1.1 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cometbft/cometbft-db v0.12.0 // indirect
	github.com/cometbft/cometbft/api v1.0.0-rc.1.0.20240806094948-2c4293ef36c4 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.13.0 // indirect
	github.com/cosmos/cosmos-db v1.0.2 // indirect
	github.com/cosmos/crypto v0.1.2 // indirect
	github.com/cosmos/gogoproto v1.5.0 // indirect
	github.com/cosmos/iavl v1.2.1-0.20240725141113-7adc688cf179 // indirect
	github.com/cosmos/ics23/go v0.10.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/dgraph-io/badger/v4 v4.2.0 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/dot v1.6.2 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.3 // indirect
	github.com/ethereum/go-ethereum v1.14.7 // indirect
	github.com/ethereum/go-verkle v0.1.1-0.20240306133620-7d920df305f0 // indirect
	github.com/ferranbt/fastssz v0.1.4-0.20240629094022-eac385e6ee79 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/getsentry/sentry-go v0.28.1 // indirect
	github.com/go-faster/xor v1.0.0 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/glog v1.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.3.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/karalabe/ssz v0.2.1-0.20240724074312-3d1ff7a6f7c4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/linxGnu/grocksdb v1.9.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/petermattis/goid v0.0.0-20240607163614-bb94eb51e7a7 // indirect
	github.com/phuslu/log v1.0.108-0.20240705160716-a8f8c12ae6c6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/prysmaticlabs/go-bitfield v0.0.0-20240618144021-706c95b2dd15 // indirect
	github.com/prysmaticlabs/gohashtree v0.0.4-beta // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/cors v1.11.0 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/supranational/blst v0.3.13 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/bbolt v1.4.0-alpha.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240711142825-46eb208f015d // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
