// Code generated by fastssz. DO NOT EDIT.
// Hash: e93affba1502d1e48f0a8ad77cdd6d95e7bdb6fa8245b30e66883b6224046d1d
// Version: 0.1.3
package state

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconStateDeneb object to a target array
func (b *BeaconStateDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Eth1GenesisHash'
	dst = append(dst, b.Eth1GenesisHash[:]...)

	// Field (1) 'RandaoMix'
	if size := len(b.RandaoMix); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconStateDeneb.RandaoMix", size, 32)
		return
	}
	dst = append(dst, b.RandaoMix...)

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconStateDeneb object
func (b *BeaconStateDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Eth1GenesisHash'
	copy(b.Eth1GenesisHash[:], buf[0:32])

	// Field (1) 'RandaoMix'
	if cap(b.RandaoMix) == 0 {
		b.RandaoMix = make([]byte, 0, len(buf[32:64]))
	}
	b.RandaoMix = append(b.RandaoMix, buf[32:64]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconStateDeneb object
func (b *BeaconStateDeneb) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconStateDeneb object with a hasher
func (b *BeaconStateDeneb) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Eth1GenesisHash'
	hh.PutBytes(b.Eth1GenesisHash[:])

	// Field (1) 'RandaoMix'
	if size := len(b.RandaoMix); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconStateDeneb.RandaoMix", size, 32)
		return
	}
	hh.PutBytes(b.RandaoMix)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconStateDeneb object
func (b *BeaconStateDeneb) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
