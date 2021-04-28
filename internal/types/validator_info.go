// Package types implements different core types of the API.
package types

import "encoding/json"

// StakerInfo holds extended staker information.
type StakerInfo struct {
	// Name represents the name of the staker
	Name *string `json:"name"`

	// LogoUrl represents staker logo URL
	LogoUrl *string `json:"logoUrl"`

	// Website represents a link to staker website
	Website *string `json:"website"`

	// Contact represents a link to contact to the staker
	Contact *string `json:"contact"`
}

// UnmarshalStakerInfo parses the JSON-encoded staker information data.
func UnmarshalStakerInfo(data []byte) (*StakerInfo, error) {
	var sti StakerInfo
	err := json.Unmarshal(data, &sti)
	return &sti, err
}

// Marshal returns the JSON encoding of staker information.
func (sti *StakerInfo) Marshal() ([]byte, error) {
	return json.Marshal(sti)
}
