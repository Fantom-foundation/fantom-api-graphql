// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// tAddress represents the reflection type of the Opera blockchain Address.
var tAddress = reflect.TypeOf(common.Address{})

// tNodeAddress represents the reflection type of the Opera network node discovery address.
var tNodeAddress = reflect.TypeOf(enode.Node{})

// AddressEncodeValue encodes Opera blockchain address into BSON data stream.
func AddressEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tAddress {
		return bsoncodec.ValueEncoderError{Name: "AddressEncodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}
	return vw.WriteString(val.Interface().(common.Address).String())
}

// AddressDecodeValue decodes Opera blockchain address from BSON data stream.
func AddressDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tAddress {
		return bsoncodec.ValueDecoderError{Name: "AddressDecodeValue", Types: []reflect.Type{tAddress}, Received: val}
	}

	var adr common.Address
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		adr = common.HexToAddress(str)
	case bsontype.Undefined:
		if err := vr.ReadUndefined(); err != nil {
			return err
		}
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(adr))
	return nil
}

// NodeAddressEncodeValue encodes Opera network node discovery address to BSON.
func NodeAddressEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tNodeAddress {
		return bsoncodec.ValueEncoderError{Name: "NodeAddressEncodeValue", Types: []reflect.Type{tNodeAddress}, Received: val}
	}

	node := val.Interface().(enode.Node)
	b, err := node.MarshalText()
	if err != nil {
		return bsoncodec.ValueEncoderError{Name: "NodeAddressEncodeValue", Types: []reflect.Type{tNodeAddress}, Received: val}
	}

	return vw.WriteString(string(b[:]))
}

// NodeAddressDecodeValue decodes Opera network node discovery address from BSON data stream.
func NodeAddressDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tNodeAddress {
		return bsoncodec.ValueDecoderError{Name: "NodeAddressDecodeValue", Types: []reflect.Type{tNodeAddress}, Received: val}
	}

	var en enode.Node
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		err = en.UnmarshalText([]byte(str))
		if err != nil {
			return err
		}
	case bsontype.Binary:
		bin, _, err := vr.ReadBinary()
		if err != nil {
			return err
		}
		err = en.UnmarshalText(bin)
		if err != nil {
			return err
		}
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(en))
	return nil
}
