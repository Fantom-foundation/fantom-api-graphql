package repository

import (
	"fantom-api-graphql/internal/logger"
	"github.com/ethereum/go-ethereum/common"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"sync"
)

// logQueueLength represents the amount of transaction logs
// allowed to be queued at a time before queue writer is slowed down
const logQueueLength = 50000

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

			/* SFC1::IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff); */
			common.HexToHash("0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9"): handleSfc1IncreasedStake,

			/* SFC1::IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff); */
			common.HexToHash("0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1"): handleSfc1IncreasedDelegation,

			/* SFC1::ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch) */
			common.HexToHash("0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce"): handleSfc1ClaimedDelegationReward,

			/* SFC1::ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch) */
			common.HexToHash("0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c"): handleSfc1ClaimedValidatorReward,

			/* SFC1::WithdrawnStake(uint256 indexed stakerID, uint256 penalty) */
			common.HexToHash("0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6"): handleSfc1WithdrawnStake,

			/* SFC1::WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty) */
			common.HexToHash("0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f"): handleSfc1WithdrawnDelegation,

			/* SFC1::DeactivatedStake(uint256 indexed stakerID) */
			common.HexToHash("0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60"): handleSfc1DeactivatedStake,

			/* SFC1::CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount) */
			common.HexToHash("0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d"): handleSfc1CreatedWithdrawRequest,

			/* SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID) */
			common.HexToHash("0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2"): handleSfc1IncreasedDelegation,

			/* SFC1::DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID) */
			common.HexToHash("0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469"): handleSfc1IncreasedDelegation,

			/* SFC1::PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty) */
			common.HexToHash("0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd"): handleSfc1PartialWithdrawByRequest,

			/* SFC1::UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount) */
			common.HexToHash("0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff"): handleSfc1UpdatedDelegation,

			/* SFC1::UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe) */
			common.HexToHash("0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c"): handleSfc1UpdatedStake,

			/* SFC3::Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount) */
			common.HexToHash("0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb"): handleSfcCreatedDelegation,

			/* SFC3::Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount) */
			common.HexToHash("0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57"): handleSfcUndelegated,

			/* SFC3::Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount) */
			common.HexToHash("0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21"): handleSfcWithdrawn,

			/* SFC3:: ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward) */
			common.HexToHash("0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45"): handleSfcClaimedRewards,

			/* SFC3::RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward) */
			common.HexToHash("0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26"): handleSfcRestakeRewards,

			/* ERC20::Approval(address indexed owner, address indexed spender, uint256 value) */
			common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"): handleErc20Approval,

			/* ERC20::Transfer(address indexed from, address indexed to, uint256 value) */
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): handleErc20Transfer,
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
