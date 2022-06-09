package resolvers

import (
	"fantom-api-graphql/internal/types"
)

const (
	ErcTrxTypeNameTransfer = "TRANSFER"
	ErcTrxTypeNameMint     = "MINT"
	ErcTrxTypeNameBurn     = "BURN"
	ErcTrxTypeNameApproval = "APPROVAL"
)

func ercTrxTypeToName(trxType int32) string {
	switch trxType {
	case types.TokenTrxTypeTransfer:
		return ErcTrxTypeNameTransfer
	case types.TokenTrxTypeApproval:
		return ErcTrxTypeNameApproval
	case types.TokenTrxTypeMint:
		return ErcTrxTypeNameMint
	case types.TokenTrxTypeBurn:
		return ErcTrxTypeNameBurn
	default:
		return "OTHER"
	}
}

func ercTrxTypesFromNames(names *[]string) []int32 {
	if names == nil {
		return nil
	}
	var vals []int32
	for _, name := range *names {
		switch name {
		case ErcTrxTypeNameTransfer:
			vals = append(vals, types.TokenTrxTypeTransfer)
		case ErcTrxTypeNameApproval:
			vals = append(vals, types.TokenTrxTypeApproval)
		case ErcTrxTypeNameMint:
			vals = append(vals, types.TokenTrxTypeMint)
		case ErcTrxTypeNameBurn:
			vals = append(vals, types.TokenTrxTypeBurn)
		}
	}
	return vals
}
