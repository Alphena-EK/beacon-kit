// Code generated by fastssz. DO NOT EDIT.
// Hash: 2a453b696b1f21668f1259864913581e35b172a4c9df1efd63d1661eacfc09b6
// Version: 0.1.3
package types

import (
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlobSidecar object
func (b *BlobSidecar) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlobSidecar object to a target array
func (b *BlobSidecar) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, b.Index)

	// Field (1) 'Blob'
	dst = append(dst, b.Blob[:]...)

	// Field (2) 'KzgCommitment'
	dst = append(dst, b.KzgCommitment[:]...)

	// Field (3) 'KzgProof'
	dst = append(dst, b.KzgProof[:]...)

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(types.BeaconBlockHeader)
	}
	if dst, err = b.BeaconBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'InclusionProof'
	if size := len(b.InclusionProof); size != 8 {
		err = ssz.ErrVectorLengthFn("BlobSidecar.InclusionProof", size, 8)
		return
	}
	for ii := 0; ii < 8; ii++ {
		dst = append(dst, b.InclusionProof[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlobSidecar object
func (b *BlobSidecar) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 131544 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	b.Index = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Blob'
	copy(b.Blob[:], buf[8:131080])

	// Field (2) 'KzgCommitment'
	copy(b.KzgCommitment[:], buf[131080:131128])

	// Field (3) 'KzgProof'
	copy(b.KzgProof[:], buf[131128:131176])

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(types.BeaconBlockHeader)
	}
	if err = b.BeaconBlockHeader.UnmarshalSSZ(buf[131176:131288]); err != nil {
		return err
	}

	// Field (5) 'InclusionProof'
	b.InclusionProof = make([][32]byte, 8)
	for ii := 0; ii < 8; ii++ {
		copy(b.InclusionProof[ii][:], buf[131288:131544][ii*32:(ii+1)*32])
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlobSidecar object
func (b *BlobSidecar) SizeSSZ() (size int) {
	size = 131544
	return
}

// HashTreeRoot ssz hashes the BlobSidecar object
func (b *BlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlobSidecar object with a hasher
func (b *BlobSidecar) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(b.Index)

	// Field (1) 'Blob'
	hh.PutBytes(b.Blob[:])

	// Field (2) 'KzgCommitment'
	hh.PutBytes(b.KzgCommitment[:])

	// Field (3) 'KzgProof'
	hh.PutBytes(b.KzgProof[:])

	// Field (4) 'BeaconBlockHeader'
	if b.BeaconBlockHeader == nil {
		b.BeaconBlockHeader = new(types.BeaconBlockHeader)
	}
	if err = b.BeaconBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'InclusionProof'
	{
		if size := len(b.InclusionProof); size != 8 {
			err = ssz.ErrVectorLengthFn("BlobSidecar.InclusionProof", size, 8)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.InclusionProof {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlobSidecar object
func (b *BlobSidecar) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
