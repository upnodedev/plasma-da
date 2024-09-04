package op_alt_da

import (
	"alt-da/common"
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
)

// CommitmentType is the commitment type prefix.
type CommitmentType byte

// CommitmentType describes the binary format of the commitment.
// KeccakCommitmentType is the default commitment type for the centralized DA storage.
// GenericCommitmentType indicates an opaque bytestring that the op-node never opens.
const (
	Keccak256CommitmentType CommitmentType = 0
	GenericCommitmentType   CommitmentType = 1
)

// CommitmentData is the binary representation of a commitment.
type CommitmentData interface {
	CommitmentType() CommitmentType
	Encode() []byte
	TxData() []byte
	Verify(input []byte) error
}

// Keccak256Commitment is an implementation of CommitmentData that uses Keccak256 as the commitment function.
type Keccak256Commitment []byte

// GenericCommitment is an implementation of CommitmentData that treats the commitment as an opaque bytestring.
type GenericCommitment []byte

// NewCommitmentData creates a new commitment from the given input and desired type.
func NewCommitmentData(t CommitmentType, input []byte) CommitmentData {
	switch t {
	case Keccak256CommitmentType:
		return NewKeccak256Commitment(input)
	case GenericCommitmentType:
		return NewGenericCommitment(input)
	default:
		return nil
	}
}

// DecodeCommitmentData parses the commitment into a known commitment type.
// The input type is determined by the first byte of the raw data.
// The input type is discarded and the commitment is passed to the appropriate constructor.
func DecodeCommitmentData(input []byte) (CommitmentData, error) {
	if len(input) == 0 {
		return nil, common.ErrInvalidCommitment
	}
	t := CommitmentType(input[0])
	data := input[1:]
	switch t {
	case Keccak256CommitmentType:
		return DecodeKeccak256(data)
	case GenericCommitmentType:
		return DecodeGenericCommitment(data)
	default:
		return nil, common.ErrInvalidCommitment
	}
}

// NewKeccak256Commitment creates a new commitment from the given input.
func NewKeccak256Commitment(input []byte) Keccak256Commitment {
	return Keccak256Commitment(crypto.Keccak256(input))
}

// DecodeKeccak256 validates and casts the commitment into a Keccak256Commitment.
func DecodeKeccak256(commitment []byte) (Keccak256Commitment, error) {
	// guard against empty commitments
	if len(commitment) == 0 {
		return nil, common.ErrInvalidCommitment
	}
	// keccak commitments are always 32 bytes
	if len(commitment) != 32 {
		return nil, common.ErrInvalidCommitment
	}
	return commitment, nil
}

// CommitmentType returns the commitment type of Keccak256.
func (c Keccak256Commitment) CommitmentType() CommitmentType {
	return Keccak256CommitmentType
}

// Encode adds a commitment type prefix self describing the commitment.
func (c Keccak256Commitment) Encode() []byte {
	return append([]byte{byte(Keccak256CommitmentType)}, c...)
}

// TxData adds an extra version byte to signal it's a commitment.
func (c Keccak256Commitment) TxData() []byte {
	return append([]byte{common.TxDataVersion1}, c.Encode()...)
}

// Verify checks if the commitment matches the given input.
func (c Keccak256Commitment) Verify(input []byte) error {
	if !bytes.Equal(c, crypto.Keccak256(input)) {
		return common.ErrCommitmentMismatch
	}
	return nil
}

// NewGenericCommitment creates a new commitment from the given input.
func NewGenericCommitment(input []byte) GenericCommitment {
	return GenericCommitment(input)
}

// DecodeGenericCommitment validates and casts the commitment into a GenericCommitment.
func DecodeGenericCommitment(commitment []byte) (GenericCommitment, error) {
	if len(commitment) == 0 {
		return nil, common.ErrInvalidCommitment
	}
	return commitment[:], nil
}

// CommitmentType returns the commitment type of Generic Commitment.
func (c GenericCommitment) CommitmentType() CommitmentType {
	return GenericCommitmentType
}

// Encode adds a commitment type prefix self describing the commitment.
func (c GenericCommitment) Encode() []byte {
	return append([]byte{byte(GenericCommitmentType)}, c...)
}

// TxData adds an extra version byte to signal it's a commitment.
func (c GenericCommitment) TxData() []byte {
	return append([]byte{common.TxDataVersion1}, c.Encode()...)
}

// Verify always returns true for GenericCommitment because the DA Server must validate the data before returning it to the op-node.
func (c GenericCommitment) Verify(input []byte) error {
	return nil
}
