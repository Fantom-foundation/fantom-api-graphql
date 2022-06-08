// Package registry provides project specific BSON encoders and decoders.
package registry

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

// defaultRegistry is the default bsoncodec.Registry. It contains the custom codecs, the default codecs and the
// primitive codecs.
var defaultRegistry = New()

// DefaultRegistry returns default bsoncodec.Registry
func DefaultRegistry() *bsoncodec.Registry {
	return defaultRegistry
}

// New creates a new BSON registry to be used for BSON marshalling/unmarshalling operations
func New() *bsoncodec.Registry {
	rb := bsoncodec.NewRegistryBuilder()

	// add defaults
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)

	// add custom codecs
	custom(rb)

	bson.PrimitiveCodecs{}.RegisterPrimitiveCodecs(rb)
	return rb.Build()
}

// custom adds project specific BSON codecs to the provided registry builder.
func custom(rb *bsoncodec.RegistryBuilder) {
	// add common.Address (value) support to the BSON registry
	rb.RegisterTypeEncoder(tAddress, bsoncodec.ValueEncoderFunc(AddressEncodeValue))
	rb.RegisterTypeDecoder(tAddress, bsoncodec.ValueDecoderFunc(AddressDecodeValue))

	// add common.Hash (value) support to the BSON registry
	rb.RegisterTypeEncoder(tHash, bsoncodec.ValueEncoderFunc(HashEncodeValue))
	rb.RegisterTypeDecoder(tHash, bsoncodec.ValueDecoderFunc(HashDecodeValue))

	// add hexutil.Big (value) support to the BSON registry
	rb.RegisterTypeEncoder(tHexBigInt, bsoncodec.ValueEncoderFunc(HexBigIntEncodeValue))
	rb.RegisterTypeDecoder(tHexBigInt, bsoncodec.ValueDecoderFunc(HexBigIntDecodeValue))

	// add hexutil.Uint variants (value) support to the BSON registry
	rb.RegisterTypeEncoder(tHexUint, bsoncodec.ValueEncoderFunc(HexUintEncodeValue))
	rb.RegisterTypeDecoder(tHexUint, bsoncodec.ValueDecoderFunc(HexUintDecodeValue))
	rb.RegisterTypeEncoder(tHexUint64, bsoncodec.ValueEncoderFunc(HexUintEncodeValue))
	rb.RegisterTypeDecoder(tHexUint64, bsoncodec.ValueDecoderFunc(HexUintDecodeValue))
}
