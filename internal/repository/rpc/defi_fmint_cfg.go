/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

//go:generate tools/abigen.sh --abi ./contracts/abi/defi-fmint-minter.abi --pkg contracts --type DefiFMintMinter --out ./contracts/fmint_minter.go

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/singleflight"
	"sync"
)

// identifiers of fMint contracts we may want to get
const (
	fMintAddressMinter             = "fantom_mint"
	fMintAddressPriceOracleProxy   = "price_oracle_proxy"
	fMintAddressRewardDistribution = "reward_distribution"
	fMintAddressTokenRegistry      = "token_registry"
	fMintCollateralPool            = "collateral_pool"
	fMintDebtPool                  = "debt_pool"
)

// fMintConfig represents the configuration for DeFi fMint module.
type fMintConfig struct {
	// bridge represents the reference to the instantiated RPC bridge
	bridge *FtmBridge

	// AddressProvider represents the address
	// of the fMint Address Provider contract
	addressProvider common.Address

	// contracts represents lazy-loaded addresses
	contracts sync.Map

	// use request group to handle contracts address resolution
	requestGroup singleflight.Group
}

// tokenRegistryContract returns an instance of the fMint TokenRegistry
// smart contract used to keep track of tokens available on fMint protocol.
func (fmc *fMintConfig) tokenRegistryContract() (*contracts.DefiFMintTokenRegistry, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressTokenRegistry)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := contracts.NewDefiFMintTokenRegistry(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint TokenRegistry contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintMinterContract returns an instance of the fMint Minter smart contract.
func (fmc *fMintConfig) fMintMinterContract() (*contracts.DefiFMintMinter, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressMinter)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := contracts.NewDefiFMintMinter(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint Minter contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintRewardsDistribution returns an instance of the fMint Reward Distribution smart contract.
func (fmc *fMintConfig) fMintRewardsDistribution() (*contracts.FMintRewardsDistribution, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressRewardDistribution)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := contracts.NewFMintRewardsDistribution(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint Rewards Distribution contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintCollateralPool returns an instance of the fMint collateral pool contract.
func (fmc *fMintConfig) fMintTokenStorage(addr common.Address) (*contracts.DeFiTokenStorage, error) {
	// connect the contract
	contract, err := contracts.NewDeFiTokenStorage(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint token pool %s; %s", addr.String(), err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintCollateralPool returns an instance of the fMint collateral pool contract.
func (fmc *fMintConfig) fMintCollateralPool() (*contracts.DeFiTokenStorage, error) {
	// get address
	addr, err := fmc.contractAddress(fMintCollateralPool)
	if err != nil {
		return nil, err
	}

	return fmc.fMintTokenStorage(addr)
}

// fMintDebtPool returns an instance of the fMint debt pool contract.
func (fmc *fMintConfig) fMintDebtPool() (*contracts.DeFiTokenStorage, error) {
	// get address
	addr, err := fmc.contractAddress(fMintDebtPool)
	if err != nil {
		return nil, err
	}

	return fmc.fMintTokenStorage(addr)
}

// priceOracleProxyContract returns an instance of the DeFi price oracle proxy
// interface smart contract.
func (fmc *fMintConfig) priceOracleProxyContract() (*contracts.PriceOracleProxyInterface, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressPriceOracleProxy)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := contracts.NewPriceOracleProxyInterface(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access DeFi PriceOracleProxy contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// mustContractAddress returns the given contract address, or an empty
// address if the real address is not available.
func (fmc *fMintConfig) mustContractAddress(name string) common.Address {
	addr, err := fmc.contractAddress(name)
	if err != nil {
		return common.Address{}
	}
	return addr
}

// contractAddress returns an address of a contract from the registry
func (fmc *fMintConfig) contractAddress(name string) (common.Address, error) {
	// make sure the contract addresses map exists
	adr, ok := fmc.contracts.Load(name)
	if ok {
		return adr.(common.Address), nil
	}

	// lazy load the contract address requested
	adr, err, _ := fmc.requestGroup.Do(name, func() (interface{}, error) {
		return fmc.pullContractAddress(name)
	})

	// handle error response
	if err != nil {
		return common.Address{}, err
	}

	// return the address
	return adr.(common.Address), nil
}

// pullContractAddress extracts the address of given fMint contract
func (fmc *fMintConfig) pullContractAddress(name string) (common.Address, error) {
	// try to get the address
	adr, err := fmc.loadAddress(name)
	if err != nil {
		return common.Address{}, err
	}

	// keep the address for later
	fmc.contracts.Store(name, *adr)
	return *adr, nil
}

// loadAddress loads a specified contract address from the AddressProvider.
func (fmc *fMintConfig) loadAddress(name string) (*common.Address, error) {
	// connect the Address Provider
	ap, err := contracts.NewDefiFMintAddressProvider(fmc.addressProvider, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint AddressProvider contract; %s", err.Error())
		return nil, err
	}

	// make the id
	var id [32]byte
	copy(id[:], name)

	// try to get the address
	addr, err := ap.GetAddress(nil, id)
	if err != nil {
		fmc.bridge.log.Errorf("[%s] can not get address of %s; %s", fmc.addressProvider.String(), name, err.Error())
		return nil, err
	}

	return &addr, nil
}
