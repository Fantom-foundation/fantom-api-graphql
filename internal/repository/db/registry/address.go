// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// tAddress represents the reflection type of the Opera blockchain Address.
var tAddress = reflect.TypeOf(common.Address{})

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
