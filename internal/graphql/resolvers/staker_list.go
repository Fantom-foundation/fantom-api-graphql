package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sort"
)

// Stakers resolves a list of staker information from SFC smart contract.
func (rs *rootResolver) Stakers() ([]*Staker, error) {
	// get the number
	num, err := repository.R().LastValidatorId()
	if err != nil {
		log.Errorf("can not get the highest staker id; %s", err.Error())
		return nil, err
	}

	// make the list
	list := make([]*Staker, 0)
	for i := uint64(1); i <= num; i++ {
		// extract the staker info
		st, err := repository.R().Validator((*hexutil.Big)(new(big.Int).SetUint64(i)))
		if err != nil {
			log.Criticalf("can not extract staker #%d information; %s", i, err.Error())
			continue
		}

		// staker not valid?
		if st.Id.ToInt().Uint64() == 0 {
			log.Debugf("staker #%d has invalid ID", i)
			continue
		}

		// store the staker into the list
		list = append(list, NewStaker(st))
	}

	// inform
	log.Debugf("found %d stakers", len(list))

	// sort the list by total amount delegated and return the result
	sort.Sort(StakesByTotalStaked(list))
	return list, nil
}

// StakesByTotalStaked represents a list of staking sortable by their total staked amount.
type StakesByTotalStaked []*Staker

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
