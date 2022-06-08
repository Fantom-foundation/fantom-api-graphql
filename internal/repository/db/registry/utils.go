// Package registry provides project specific BSON encoders and decoders.
package registry

import "go.mongodb.org/mongo-driver/bson"

// Marshal performed with custom registry
func Marshal(val interface{}) ([]byte, error) {
	return bson.MarshalWithRegistry(DefaultRegistry(), val)
}

// Unmarshal performed with custom registry
func Unmarshal(data []byte, v interface{}) error {
	return bson.UnmarshalWithRegistry(DefaultRegistry(), data, v)
}
