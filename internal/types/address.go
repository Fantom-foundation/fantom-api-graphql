package types

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"reflect"
)

// Address represents an address of an object on chain.
type Address common.Address

// hex prefix represents the prefix expected on hex coding of unmarshaled data.
var hexPrefix = []byte{0x30, 0x78}

var (
	// addressT represents the Address type reflection
	addressT = reflect.TypeOf(Address{})
)

// MarshalBSON parses a hash in hex syntax.
func (a Address) String() string {
	return common.Address(a).String()
}

// MarshalText returns the hex representation of a.
func (a Address) MarshalText() ([]byte, error) {
	return hexutil.Bytes(a[:]).MarshalText()
}

// MarshalBSON parses a hash in hex syntax.
func (a Address) MarshalBSON() ([]byte, error) {
	return hexutil.Bytes(a[:]).MarshalText()
}

// UnmarshalJSON parses a hash in hex syntax.
func (a *Address) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(addressT, input, a[:])
}

// UnmarshalText parses a hash in hex syntax.
func (a *Address) UnmarshalBSON(input []byte) error {
	// make sure to skip any MongoDB prefix for the address
	if len(input) > 1 && input[0] != '0' {
		// find the expected prefix position
		pos := bytes.Index(input, hexPrefix)
		if pos < 0 {
			return fmt.Errorf("unknown address format")
		}

		// log and convert
		*a = Address(common.HexToAddress(string(input[pos:])))
		return nil
	}

	// any data?
	if 2 < len(input) {
		return hexutil.UnmarshalFixedText("Address", input, a[:])
	}
	return nil
}

// UnmarshalText parses a hash in hex syntax.
func (a *Address) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Address", input, a[:])
}

// ImplementsGraphQLType returns true if Hash implements the specified GraphQL type.
func (a Address) ImplementsGraphQLType(name string) bool { return name == "Address" }

// UnmarshalGraphQL un-marshals the provided GraphQL query data.
func (a *Address) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		err = a.UnmarshalText([]byte(input))
	default:
		err = fmt.Errorf("unexpected type %T for Address", input)
	}
	return err
}
