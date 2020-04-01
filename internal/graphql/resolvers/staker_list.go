package resolvers

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sort"
)

// Resolves a staker information from SFC smart contract.
func (rs *rootResolver) Stakers() ([]Staker, error) {
	// get the number
	num, err := rs.repo.LastStakerId()
	if err != nil {
		rs.log.Errorf("can not get the highest staker id; %s", err.Error())
		return nil, err
	}

	// make the list
	list := make([]Staker, num)

	// loop to get the data
	for i := uint64(0); i < uint64(num); i++ {
		// extract the staker info
		st, err := rs.repo.Staker(hexutil.Uint64(i + 1))
		if err != nil {
			rs.log.Criticalf("can not extract staker information; %s", err.Error())
		} else {
			// store to the list
			list[i] = *NewStaker(st, rs.repo)
		}
	}

	// sort the list by total amount delegated and return the result
	sort.Sort(StakesByTotalStaked(list))
	return list, nil
}

// StakesByTotalStaked represents a list of staking sortable by their total staked amount.
type StakesByTotalStaked []Staker

// Len returns size of the stakers list.
func (s StakesByTotalStaked) Len() int {
	return len(s)
}

// Less compares two stakers and returns true if the first is lower than the last.
// We use it to sort stakers by the total amount on stake.
func (s StakesByTotalStaked) Less(i, j int) bool {
	return s[i].TotalStake.ToInt().Cmp(s[j].TotalStake.ToInt()) > 0
}

// Swap changes position of two stakers in the list.
func (s StakesByTotalStaked) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
