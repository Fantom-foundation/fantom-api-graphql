// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

// tHash represents the reflection type of the Opera blockchain generic Hash.
var tHash = reflect.TypeOf(common.Hash{})

// HashEncodeValue encodes Opera blockchain hash into BSON data stream.
func HashEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tHash {
		return bsoncodec.ValueEncoderError{Name: "HashEncodeValue", Types: []reflect.Type{tHash}, Received: val}
	}
	return vw.WriteString(val.Interface().(common.Hash).String())
}

// HashDecodeValue decodes Opera blockchain hash from BSON data stream.
func HashDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHash {
		return bsoncodec.ValueDecoderError{Name: "HashDecodeValue", Types: []reflect.Type{tHash}, Received: val}
	}

	var hash common.Hash
	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		hash = common.HexToHash(str)
	default:
		return vr.ReadUndefined()
	}

	val.Set(reflect.ValueOf(hash))
	return nil
}
