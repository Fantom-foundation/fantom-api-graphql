// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"reflect"
)

// tHexBytes represents the reflection type of the hexutil.Bytes array.
var tHexBytes = reflect.TypeOf(hexutil.Bytes{})

// HexBytesEncodeValue encodes hexutil.Bytes into BSON data stream.
func HexBytesEncodeValue(con bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tHexBytes {
		return bsoncodec.ValueEncoderError{Name: "HexBytesEncodeValue", Types: []reflect.Type{tHexBytes}, Received: val}
	}

	return bsoncodec.NewByteSliceCodec().EncodeValue(con, vw, reflect.ValueOf([]byte(val.Interface().(hexutil.Bytes))))
}

// HexBytesDecodeValue decodes hexutil.Bytes from BSON data stream.
func HexBytesDecodeValue(con bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHexBytes {
		return bsoncodec.ValueDecoderError{Name: "HexBytesDecodeValue", Types: []reflect.Type{tHexBytes}, Received: val}
	}

	err := bsoncodec.NewByteSliceCodec().DecodeValue(con, vr, val)
	if err != nil {
		return err
	}

	val.Set(reflect.ValueOf(hexutil.Bytes(val.Interface().([]byte))))

	return nil
}
