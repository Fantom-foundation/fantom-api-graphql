package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RewardClaim represents resolvable reward claim detail.
type RewardClaim struct {
	types.RewardClaim
}

// NewRewardClaim creates new instance of resolvable reward claim.
func NewRewardClaim(c *types.RewardClaim) *RewardClaim {
	return &RewardClaim{RewardClaim: *c}
}

// Address resolves the address of the delegation.
func (rwc RewardClaim) Address() common.Address {
	return rwc.Delegator
}

// ToStakerId resolves the ID of the validator of the claim.
func (rwc RewardClaim) ToStakerId() hexutil.Big {
	return rwc.ToValidatorId
}

// IsRestaked resolves the restake mark of the claim.
func (rwc RewardClaim) IsRestaked() bool {
	return rwc.IsDelegated
}

// TrxHash resolves the hash of the claim transaction.
func (rwc RewardClaim) TrxHash() common.Hash {
	return rwc.ClaimTrx
}
