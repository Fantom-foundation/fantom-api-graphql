// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
)

var (
	tHexUint        = reflect.TypeOf(hexutil.Uint(0))
	tHexUint64      = reflect.TypeOf(hexutil.Uint64(0))
	defaultEncoders = bsoncodec.DefaultValueEncoders{}
)

// HexUintEncodeValue encodes hexutil.Uint and/or hexutil.Uint64 into BSON data stream.
func HexUintEncodeValue(con bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	switch val.Kind() {
	case tHexUint.Kind():
		v := val.Interface().(hexutil.Uint)
		intVal := int64(v)
		// encode as standard int if fits
		if intVal >= 0 {
			return defaultEncoders.IntEncodeValue(con, vw, reflect.ValueOf(intVal))
		}
		return vw.WriteString((&v).String())

	case tHexUint64.Kind():
		v := val.Interface().(hexutil.Uint64)
		intVal := int64(v)
		// encode as standard int if fits
		if intVal >= 0 {
			return defaultEncoders.IntEncodeValue(con, vw, reflect.ValueOf(intVal))
		}
		return vw.WriteString((&v).String())
	}

	return bsoncodec.ValueEncoderError{Name: "HexUintEncodeValue", Types: []reflect.Type{tHexUint, tHexUint64}, Received: val}
}

// HexUintDecodeValue decodes hexutil.Uint and/or hexutil.Uint64 from BSON data stream.
func HexUintDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tHexUint {
		return bsoncodec.ValueDecoderError{Name: "HexUintDecodeValue", Types: []reflect.Type{tHexUint}, Received: val}
	}

	var value uint64

	switch vrType := vr.Type(); vrType {
	case bsontype.String:
		str, err := vr.ReadString()
		if err != nil {
			return err
		}
		value, err = hexutil.DecodeUint64(str)
		if err != nil {
			return err
		}
	case bsontype.Int32:
		i32, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		value = uint64(i32)
	case bsontype.Int64:
		i64, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		value = uint64(i64)
	default:
		return vr.ReadUndefined()
	}

	switch val.Kind() {
	case tHexUint.Kind():
		val.Set(reflect.ValueOf(hexutil.Uint(value)))
	case tHexUint64.Kind():
		val.Set(reflect.ValueOf(hexutil.Uint64(value)))
	default:
		return bsoncodec.ValueDecoderError{
			Name:     "HexUintDecodeValue",
			Kinds:    []reflect.Kind{tHexUint.Kind(), tHexUint64.Kind()},
			Received: reflect.Zero(val.Type()),
		}
	}
	return nil
}
