/*
Package rpc implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

//go:generate abigen --abi ./contracts/ballot.abi --pkg rpc --type BallotContract --out ./ballot_bind.go

// ballotDetails is the structure describing basic ballot details stored in the contract
type ballotDetails struct {
	Name      [32]byte
	Url       string
	Start     *big.Int
	End       *big.Int
	Finalized bool
}

// VotesList represents a list of votes.
type VotesList []types.Vote

// LoadBallotDetails loads additional ballot details from a deployed smart contract.
// The initial ballot address is supposed to be already available
// so the contract connection can be made.
func (ftm *FtmBridge) LoadBallotDetails(ballot *types.Ballot) error {
	// connect the contract
	contract, err := NewBallotContract(ballot.Address, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open ballot contract connection; %s", err.Error())
		return err
	}

	// get details if possible
	details, err := ftm.ballotDetails(contract)
	if err != nil {
		ftm.log.Errorf("ballot details not available; %s", err.Error())
		return err
	}

	// extend the ballot with loaded details
	extendBallotFromDetails(ballot, details)

	// load the list of proposals
	ballot.Proposals, err = ftm.ballotProposals(contract)
	if err != nil {
		ftm.log.Errorf("ballot details not available; %s", err.Error())
		return err
	}

	return nil
}

// BallotIsFinalized checks if the given ballot has already been finalized.
func (ftm *FtmBridge) BallotIsFinalized(addr *common.Address) (bool, error) {
	// address has to be valid
	if addr == nil {
		return false, fmt.Errorf("missing ballot address")
	}

	// connect the contract
	contract, err := NewBallotContract(*addr, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open ballot contract connection; %s", err.Error())
		return false, err
	}

	// read ballot details if possible
	details, err := ftm.ballotDetails(contract)
	if err != nil {
		ftm.log.Errorf("ballot details not available; %s", err.Error())
		return false, err
	}

	return details.Finalized, nil
}

// BallotWinner checks if the given ballot has already been finalized and has a winner.
// The function returns nil if the winner is not decided yet.
func (ftm *FtmBridge) BallotWinner(addr *common.Address) (*hexutil.Uint64, error) {
	// address has to be valid
	if addr == nil {
		return nil, fmt.Errorf("missing ballot address")
	}

	// connect the contract
	contract, err := NewBallotContract(*addr, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open ballot contract connection; %s", err.Error())
		return nil, err
	}

	// try to get the winner
	win, _, _, err := contract.Winner(nil)
	if err != nil {
		ftm.log.Errorf("can not read ballot winner information; %s", err.Error())
		return nil, err
	}

	// do we have the winner?
	if win == nil {
		return nil, nil
	}

	// convert the winning index
	index := hexutil.Uint64(win.Uint64())
	return &index, nil
}

// Votes exports list of votes for the given voter on a list of ballots.
func (ftm *FtmBridge) Votes(voter common.Address, ballots []common.Address) (VotesList, error) {
	// make the container
	list := make([]types.Vote, 0)

	// loop all ballots in question
	for _, ballot := range ballots {
		// pull the votes for this ballot
		sub, err := ftm.VotesByBallot(ballot, &voter)
		if err != nil {
			return nil, err
		}

		// add to the final list
		list = append(list, sub...)
	}

	return list, nil
}

// VotesByBallot export list of all votes on the given ballot.
// Optionally, the voter can be filtered by the wallet address.
func (ftm *FtmBridge) VotesByBallot(ballot common.Address, voter *common.Address) (VotesList, error) {
	// get the iterator
	it, err := ftm.votesIterator(ballot, voter)
	if err != nil {
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			ftm.log.Errorf("failed to close votes iterator; %s", err.Error())
		}
	}()

	// make the container
	list := make(VotesList, 0)

	// pull the list from the initialized iterator
	err = ftm.votesList(it, &list)
	if err != nil {
		return nil, err
	}

	ftm.log.Infof("Ballot %s, votes %d", ballot.String(), len(list))

	return list, nil
}

//  votesIterator creates a new vote iterator for the given ballot and optional voter.
func (ftm *FtmBridge) votesIterator(ballot common.Address, voter *common.Address) (*BallotContractVotedIterator, error) {
	// connect the contract
	contract, err := NewBallotContract(ballot, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open ballot contract connection; %s", err.Error())
		return nil, err
	}

	// prep iteration variables
	var it *BallotContractVotedIterator

	// create event iterator for the votes
	if voter == nil {
		// we don't care about the voter address
		it, err = contract.FilterVoted(nil, []common.Address{ballot}, nil)
	} else {
		// we want to pull votes only for specified voter address
		it, err = contract.FilterVoted(nil, []common.Address{ballot}, []common.Address{*voter})
	}

	// check for errors before will try to pull
	if err != nil {
		ftm.log.Errorf("failed to get filtered votes iterator; %s", err.Error())
		return nil, err
	}

	return it, nil
}

// votedByBallotList creates a list of votes from the given filtered iterator.
func (ftm *FtmBridge) votesList(it *BallotContractVotedIterator, list *VotesList) error {
	// loop through the iterator
	for it.Next() {
		// make sure this is a valid record
		if it.Event.Vote == nil {
			ftm.log.Error("invalid vote record found")
			continue
		}

		// get the vote index
		ix := hexutil.Uint64(it.Event.Vote.Uint64())

		// populate the local struct with data we need
		vote := types.Vote{
			Voter:  it.Event.Voter,
			Ballot: it.Event.Ballot,
			Vote:   &ix,
		}

		// add to the list
		inList := *list
		*list = append(inList, vote)
	}

	// detect possible error in traversing the iterator
	if err := it.Error(); err != nil {
		ftm.log.Errorf("failed to pull from filtered votes iterator; %s", err.Error())
		return err
	}

	return nil
}

// extendBallotFromDetails transfers information from internal details
// structure into the general Ballot type.
func extendBallotFromDetails(ballot *types.Ballot, details *ballotDetails) {
	// transfer the basics
	ballot.Name = strings.TrimRight(string(details.Name[:]), "\u0000 ")

	// decode URL address
	url, err := hexutil.Decode(details.Url)
	if err == nil {
		ballot.DetailsUrl = strings.TrimRight(string(url[:]), "\u0000 ")
	}

	// start date/time
	if details.Start != nil {
		ballot.Start = hexutil.Uint64(details.Start.Uint64())
	}

	// end date/time
	if details.End != nil {
		ballot.End = hexutil.Uint64(details.End.Uint64())
	}
}

// ballotDetails loads additional ballot details from a deployed smart contract.
// The initial ballot address is supposed to be already available
// so the contract connection can be made.
func (ftm *FtmBridge) ballotDetails(contract *BallotContract) (*ballotDetails, error) {
	// get base ballot details
	details, err := contract.Ballot(nil)
	if err != nil {
		ftm.log.Errorf("can not read ballot details; %s", err.Error())
		return nil, err
	}

	// make the conversion
	det := ballotDetails(details)
	return &det, nil
}

// ballotProposals load list of ballot proposal names from the smart contract.
func (ftm *FtmBridge) ballotProposals(contract *BallotContract) ([]string, error) {
	// get the number of proposals in the ballot
	count, err := contract.ProposalsCount(nil)
	if err != nil {
		ftm.log.Errorf("can not read ballot proposals range; %s", err.Error())
		return nil, err
	}

	// make a container for proposals
	list := make([]string, 0)
	index := new(big.Int)

	// load all the proposals
	for i := uint64(0); i < count.Uint64(); i++ {
		// read the proposal
		prop, err := contract.Proposals(nil, index.SetUint64(i))
		if err != nil {
			ftm.log.Errorf("can not read proposal %d from the contract; %s", i, err.Error())
			return nil, err
		}

		// add the proposal to the list
		list = append(list, strings.TrimRight(string(prop.Name[:]), "\u0000 "))
	}

	return list, nil
}
