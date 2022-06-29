// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/p2p/enode"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// tNodeAddress represents the reflection type of the Opera network node as defined by EIP-778.
var tNodeAddress = reflect.TypeOf(enode.Node{})

// tNodeID represents the reflection type of the Opera network node ID.
var tNodeID = reflect.TypeOf(enode.ID{})

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

// NodeIDEncodeValue encodes Opera network node ID to BSON.
func NodeIDEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tNodeID {
		return bsoncodec.ValueEncoderError{Name: "NodeIDEncodeValue", Types: []reflect.Type{tNodeID}, Received: val}
	}

	node := val.Interface().(enode.ID)
	b, err := node.MarshalText()
	if err != nil {
		return bsoncodec.ValueEncoderError{Name: "NodeIDEncodeValue", Types: []reflect.Type{tNodeID}, Received: val}
	}

	return vw.WriteString(string(b[:]))
}

// NodeIDDecodeValue decodes Opera network node ID from BSON data stream.
func NodeIDDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tNodeID {
		return bsoncodec.ValueDecoderError{Name: "NodeIDDecodeValue", Types: []reflect.Type{tNodeID}, Received: val}
	}

	var en enode.ID
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
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(en))
	return nil
}
