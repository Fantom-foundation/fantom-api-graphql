// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// logDispatcher implements dispatcher of new log events in the blockchain.
type logDispatcher struct {
	service
	inLog       chan *types.LogRecord
	knownTopics map[common.Hash]func(*types.LogRecord)
}

// name returns the name of the service used by orchestrator.
func (lgd *logDispatcher) name() string {
	return "log dispatcher"
}

// init prepares the log dispatcher to perform its function.
func (lgd *logDispatcher) init() {
	lgd.sigStop = make(chan struct{})
	lgd.knownTopics = map[common.Hash]func(*types.LogRecord){
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

		/* SFC1::UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards) */
		common.HexToHash("0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126"): handleSfc1UnstashedReward,

		/* SFC1::DeactivatedStake(uint256 indexed stakerID) */
		common.HexToHash("0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60"): handleSfc1DeactivatedStake,

		/* SFC1::PreparedToWithdrawStake(uint256 indexed stakerID) */
		common.HexToHash("0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185"): handleSfc1DeactivatedStake,

		/* SFC1::DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID) */
		common.HexToHash("0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469"): handleSfc1DeactivatedDelegation,

		/* SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID) */
		common.HexToHash("0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2"): handleSfc1DeactivatedDelegation,

		/* SFC1::CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount) */
		common.HexToHash("0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d"): handleSfc1CreatedWithdrawRequest,

		/* SFC1::WithdrawnStake(uint256 indexed stakerID, uint256 penalty) */
		common.HexToHash("0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6"): handleSfc1WithdrawnStake,

		/* SFC1::WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty) */
		common.HexToHash("0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f"): handleSfc1WithdrawnDelegation,

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

		/* SFC3::LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount) */
		common.HexToHash("0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78"): handleLockedUpStake,

		/* SFC3::UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty) */
		common.HexToHash("0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9"): handleUnlockedStake,

		/* ---------------- ERC20 and ERC721 contracts related event hooks below this line ---------------- */

		/* ERC20::Approval(address indexed owner, address indexed spender, uint256 value) */
		common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"): handleErcTokenApproval,

		/* ERC20::Transfer(address indexed from, address indexed to, uint256 value) */
		common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): handleErcTokenTransfer,

		/* ERC1155::TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value) */
		common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"): handleErc1155TransferSingle,

		/* ERC1155::TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values) */
		common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"): handleErc1155TransferBatch,

		/* --------------------- Uniswap contract related event hooks below this line --------------------- */

		/* UniswapPair::Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to) */
		common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"): handleUniswapSwap,

		/* UniswapPair::Mint(address indexed sender, uint256 amount0, uint256 amount1) */
		common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"): handleUniswapMint,

		/* UniswapPair::Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to) */
		common.HexToHash("0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496"): handleUniswapBurn,

		/* UniswapPair::Sync(uint112 reserve0, uint112 reserve1) */
		common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"): handleUniswapSync,

		/* ---------------------- fMint contract related event hooks below this line ----------------------- */

		/* FantomMintCollateral::Deposited(address indexed token, address indexed user, uint256 amount) */
		common.HexToHash("0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7"): handleFMintDeposit,

		/* FantomMintCollateral::Withdrawn(address indexed token, address indexed user, uint256 amount) */
		common.HexToHash("0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb"): handleFMintWithdraw,

		/* FantomMintDebt::Minted(address indexed token, address indexed user, uint256 amount, uint256 fee) */
		common.HexToHash("0x03f17d66ad3bf18e9412eb06582908831508cdb9b8da9cddb1431f645a5b8632"): handleFMintMint,

		/* FantomMintDebt::Repaid(address indexed token, address indexed user, uint256 amount) */
		common.HexToHash("0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a"): handleFMintRepay,

		/* FantomMintRewardManager::RewardPaid(address indexed user, uint256 reward) */
		common.HexToHash("0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486"): handleFMintReward,
	}
}

// run starts the transaction logs dispatcher job
func (lgd *logDispatcher) run() {
	// make sure we are orchestrated
	if lgd.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", lgd.name()))
	}

	// signal orchestrator we started and go
	lgd.mgr.started(lgd)
	go lgd.execute()
}

// execute implements the dispatcher reader and router routine.
func (lgd *logDispatcher) execute() {
	// don't forget to sign off after we are done
	defer func() {
		lgd.mgr.finished(lgd)
	}()

	// wait for logs and process them
	for {
		// try to read next transaction
		select {
		case <-lgd.sigStop:
			return
		case lr, ok := <-lgd.inLog:
			// is the channel even available for reading
			if !ok {
				log.Notice("logs channel closed, terminating %s", lgd.name())
				return
			}

			// try to find the topic handler
			if nil != lr && nil != lr.Topics && 0 < len(lr.Topics) {
				handler, ok := lgd.knownTopics[lr.Topics[0]]
				if ok && lr.Block != nil && lr.Trx != nil {
					log.Debugf("known topic %s found, processing", lr.Topics[0].String())
					handler(lr)
				}
			}

			// mark the processing of this log record as finished
			lr.WatchDog.Done()
		}
	}
}
