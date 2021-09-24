package resolvers

import (
	"fantom-api-graphql/internal/types"
)

const (
	ErcTrxTypeNameTransfer       = "TRANSFER"
	ErcTrxTypeNameMint           = "MINT"
	ErcTrxTypeNameBurn           = "BURN"
	ErcTrxTypeNameApproval       = "APPROVAL"
	ErcTrxTypeNameApprovalForAll = "APPROVAL_FOR_ALL"
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
	case types.TokenTrxTypeApprovalForAll:
		return ErcTrxTypeNameApprovalForAll
	default:
		return "OTHER"
	}
}

func ercTrxTypeFromName(name string) int32 {
	switch name {
	case ErcTrxTypeNameTransfer:
		return types.TokenTrxTypeTransfer
	case ErcTrxTypeNameApproval:
		return types.TokenTrxTypeApproval
	case ErcTrxTypeNameMint:
		return types.TokenTrxTypeMint
	case ErcTrxTypeNameBurn:
		return types.TokenTrxTypeBurn
	case ErcTrxTypeNameApprovalForAll:
		return types.TokenTrxTypeApprovalForAll
	default:
		return 0
	}
}
