package repository

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiConfiguration resolves the current DeFi contract settings.
func (p *proxy) DefiConfiguration() (*types.DefiSettings, error) {
	return p.rpc.DefiConfiguration()
}

// DefiToken loads details of a single DeFi token by it's address.
func (p *proxy) DefiToken(token *common.Address) (*types.DefiToken, error) {
	return p.rpc.DefiToken(token)
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (p *proxy) DefiTokens() ([]types.DefiToken, error) {
	return p.rpc.DefiTokens()
}

// DefiTokenPrice loads the current price of the given token
// from on-chain price oracle.
func (p *proxy) DefiTokenPrice(token *common.Address) (hexutil.Big, error) {
	return p.rpc.FMintTokenPrice(token)
}

// FMintAccount loads details of a DeFi/fMint account identified by the owner address.
func (p *proxy) FMintAccount(owner common.Address) (*types.FMintAccount, error) {
	return p.rpc.FMintAccount(&owner)
}

// FMintTokenBalance loads balance of a single DeFi token by it's address.
func (p *proxy) FMintTokenBalance(owner *common.Address, token *common.Address, tp types.DefiTokenType) (hexutil.Big, error) {
	return p.rpc.FMintTokenBalance(owner, token, tp)
}

// FMintTokenTotalBalance loads total balance of a single DeFi token by it's address.
func (p *proxy) FMintTokenTotalBalance(token *common.Address, tp types.DefiTokenType) (hexutil.Big, error) {
	return p.rpc.FMintTokenTotalBalance(token, tp)
}

// FMintTokenValue loads value of a single DeFi token by it's address in fUSD.
func (p *proxy) FMintTokenValue(owner *common.Address, token *common.Address, tp types.DefiTokenType) (hexutil.Big, error) {
	return p.rpc.FMintTokenValue(owner, token, tp)
}

// FMintRewardsEarned represents the total amount of rewards
// accumulated on the account for the excessive collateral deposits.
func (p *proxy) FMintRewardsEarned(addr *common.Address) (hexutil.Big, error) {
	return p.rpc.FMintRewardsEarned(addr)
}

// FMintRewardsStashed represents the total amount of rewards
// accumulated on the account in stash.
func (p *proxy) FMintRewardsStashed(addr *common.Address) (hexutil.Big, error) {
	return p.rpc.FMintRewardsStashed(addr)
}

// FMintCanClaimRewards resolves the fMint account flag for being allowed
// to claim earned rewards.
func (p *proxy) FMintCanClaimRewards(addr *common.Address) (bool, error) {
	return p.rpc.FMintCanClaimRewards(addr)
}

// FMintCanReceiveRewards resolves the fMint account flag for being eligible
// to receive earned rewards. If the collateral to debt ration drop below
// certain value, earned rewards are burned.
func (p *proxy) FMintCanReceiveRewards(addr *common.Address) (bool, error) {
	return p.rpc.FMintCanReceiveRewards(addr)
}

// FMintCanPushRewards signals if there are any rewards unlocked
// on the rewards distribution contract and can be pushed to accounts.
func (p *proxy) FMintCanPushRewards() (bool, error) {
	return p.rpc.FMintCanPushRewards()
}

// FLendGetLendingPool resolves lending pool contract instace
// to be able to get calls and informations from this contract
func (p *proxy) FLendGetLendingPool() (*contracts.ILendingPool, error) {
	return p.rpc.FLendGetLendingPool()
}

// FLendGetLendingPoolReserveData resolves reserve data
// according to given address
func (p *proxy) FLendGetLendingPoolReserveData(assetAddress *common.Address) (*types.ReserveData, error) {
	return p.rpc.FLendGetLendingPoolReserveData(assetAddress)
}

// FLendGetUserAccountData resolves user account data for
// specified address
func (p *proxy) FLendGetUserAccountData(userAddress *common.Address) (*types.FLendUserAccountData, error) {
	return p.rpc.FLendGetUserAccountData(userAddress)
}

// FLendGetReserveList resolves list of reserves in lending pool
func (p *proxy) FLendGetReserveList() ([]common.Address, error) {
	return p.rpc.FLendGetReserveList()
}

// FLendGetUserDepositHistory resolves deposit history
// data for specified user and asset address
func (p *proxy) FLendGetUserDepositHistory(userAddress *common.Address, assetAddress *common.Address) ([]*types.FLendDeposit, error) {
	return p.rpc.FLendGetUserDepositHistory(userAddress, assetAddress)
}
