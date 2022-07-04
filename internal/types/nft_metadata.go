package types

import "encoding/json"

// NftMetadata describes a token as defined in ERC-721/ERC-1155 Metadata JSON Schema.
type NftMetadata struct {
	Name        *string `json:"name"`        // Identifies the asset to which this token represents
	Description *string `json:"description"` // Describes the asset to which this token represents
	Image       *string `json:"image"`       // A URI pointing to a resource with mime type image/* representing the asset to which this token represents.
	Decimals    *int    `json:"decimals"`    // ERC-1155 only: The number of decimal places that the token amount should display
}

// DecodeNftMetadata parses the NFT token Metadata JSON.
func DecodeNftMetadata(data []byte) (*NftMetadata, error) {
	var out NftMetadata
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
