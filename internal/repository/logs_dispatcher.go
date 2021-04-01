package repository

import (
	"fantom-api-graphql/internal/logger"
	"github.com/ethereum/go-ethereum/common"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"sync"
)

// logQueueLength represents the amount og transaction logs
// allowed to be queued at a time before queue writer is slowed down
const logQueueLength = 10000

// eventTrxLog represents a log record to be processed.
type eventTrxLog struct {
	wg *sync.WaitGroup
	retypes.Log
}

// logsDispatcher implements dispatcher of new log events in the blockchain.
type logsDispatcher struct {
	service
	buffer      chan *eventTrxLog
	knownTopics map[common.Hash]func(*retypes.Log, *logsDispatcher)
}

// newLogsDispatcher creates a new transaction logs dispatcher instance.
func newLogsDispatcher(buffer chan *eventTrxLog, repo Repository, log logger.Logger, wg *sync.WaitGroup) *logsDispatcher {
	// create new dispatcher
	return &logsDispatcher{
		service: newService("logs dispatcher", repo, log, wg),
		buffer:  buffer,
		knownTopics: map[common.Hash]func(*retypes.Log, *logsDispatcher){
			/* SFC1::CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount) */
			common.HexToHash("0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be"): handleSfcCreatedDelegation,

			/* SFC1::CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount) */
			common.HexToHash("0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39"): handleSfcCreatedStake,

			/* SFC3::Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount) */
			common.HexToHash("0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb"): handleSfcCreatedDelegation,

			/* SFC3::Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount) */
			common.HexToHash("0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57"): handleSfcUndelegated,

			/* SFC3::Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount) */
			common.HexToHash("0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21"): handleSfcWithdrawn,

			/* SFC3::event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward); */
			common.HexToHash("0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26"): handleSfcRestake,
		},
	}
}

// run starts the transaction logs dispatcher job
func (ld *logsDispatcher) run() {
	ld.wg.Add(1)
	go ld.dispatch()
}

// dispatch implements the dispatcher reader and router routine.
func (ld *logsDispatcher) dispatch() {
	// log action
	ld.log.Notice("logs dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		ld.log.Notice("logs dispatcher is closed")
		ld.wg.Done()
	}()

	// wait for logs and process them
	for {
		// try to read next transaction
		select {
		case log := <-ld.buffer:
			// try to find the topic handler
			handler, ok := ld.knownTopics[log.Topics[0]]
			if ok {
				ld.log.Debugf("known topic %s found, processing", log.Topics[0].String())
				handler(&log.Log, ld)
			}

			// mark the processing as finished
			log.wg.Done()

		case <-ld.sigStop:
			return
		}
	}
}
