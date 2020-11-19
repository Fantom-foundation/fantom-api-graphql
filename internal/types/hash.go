// Package types implements different core types of the API.
package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"reflect"
)

// HashLength is the expected length of the hash type.
const HashLength = 32

// Hash represents the 32 byte hash of arbitrary data.
// It's inspired by go-ethereum Hash type, just doesn't follow some of the naming conventions.
type Hash [HashLength]byte

// hashT holds the reflected type of Hash.
var hashT = reflect.TypeOf(Hash{})

// Hex converts a hash to a hex string.
func (h Hash) Hex() string {
	return hexutil.Encode(h[:])
}

// String implements the stringer interface and is used also by the logger to attach Hash information.
func (h Hash) String() string {
	return h.Hex()
}

// Format implements fmt.Formatter, the byte slice to be formatted as is, without going through the stringer.
func (h Hash) Format(s fmt.State, c rune) {
	if _, err := fmt.Fprintf(s, "%"+string(c), h[:]); err != nil {
		return
	}
}

// Big converts a hash to a big integer.
func (h Hash) Big() *big.Int {
	return new(big.Int).SetBytes(h[:])
}

// UnmarshalText parses a hash in hex syntax.
func (h *Hash) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Hash", input, h[:])
}

// UnmarshalJSON parses a hash in hex syntax from JSON input.
func (h *Hash) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(hashT, input, h[:])
}

// UnmarshalText parses a hash in hex syntax.
func (h *Hash) UnmarshalBSON(input []byte) error {
	// make sure to skip any MongoDB prefix for the address
	if len(input) > 1 && input[0] != '0' {
		// find the expected prefix position
		pos := bytes.Index(input, hexPrefix)
		if pos < 0 {
			return fmt.Errorf("unknown hash format")
		}

		// log and convert
		*h = Hash(common.HexToHash(string(input[pos:])))
		return nil
	}

	// any data here
	if 2 < len(input) {
		return hexutil.UnmarshalFixedText("Hash", input, h[:])
	}

	return nil
}

// MarshalText returns the hex representation of Hash.
func (h Hash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// Scan implements Scanner interface for reading Hash from database/sql.
func (h *Hash) Scan(src interface{}) error {
	// get raw bytes
	srcB, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not scan %T into Hash", src)
	}

	// do we get the right size
	if len(srcB) != HashLength {
		return fmt.Errorf("can not scan []byte of len %d into Hash, expected %d", len(srcB), HashLength)
	}

	copy(h[:], srcB)
	return nil
}

// Value implements Valuer interface for storing Hash in a database/sql.
func (h Hash) Value() (driver.Value, error) {
	return h[:], nil
}

// UnmarshalGraphQL unmarshal the provided GraphQL query data.
func (h *Hash) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		err = h.UnmarshalText([]byte(input))
	default:
		err = fmt.Errorf("unexpected type %T for Hash", input)
	}
	return err
}

// SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

// ImplementsGraphQLType returns true if Hash implements the specified GraphQL type, in this case it implements Hash.
func (Hash) ImplementsGraphQLType(name string) bool {
	return name == "Hash"
}

// BytesToHash sets b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

// HexToHash sets byte representation of s to hash.
// If b is larger than len(h), b will be cropped from the left.
func HexToHash(s string) Hash {
	return BytesToHash(common.FromHex(s))
}
