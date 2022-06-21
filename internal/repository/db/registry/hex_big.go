// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// tHexBigInt represents the reflection type of the hexutil.Big integer.
var tHexBigInt = reflect.TypeOf(hexutil.Big{})

// HexBigIntEncodeValue encodes hexutil.Big into BSON data stream.
func HexBigIntEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tHexBigInt {
		return bsoncodec.ValueEncoderError{Name: "HexBigIntEncodeValue", Types: []reflect.Type{tHexBigInt}, Received: val}
	}
	v := val.Interface().(hexutil.Big)
	return vw.WriteString((&v).String())
}

// HexBigIntDecodeValue decodes hexutil.Big from BSON data stream.
func HexBigIntDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHexBigInt {
		return bsoncodec.ValueDecoderError{Name: "HexBigIntDecodeValue", Types: []reflect.Type{tHexBigInt}, Received: val}
	}

	var big hexutil.Big
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}

		v, err := hexutil.DecodeBig(str)
		if err != nil {
			return err
		}

		big = hexutil.Big(*v)
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(big))
	return nil
}
