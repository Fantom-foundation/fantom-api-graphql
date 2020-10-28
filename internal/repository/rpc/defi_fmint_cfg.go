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

//go:generate abigen --abi ./contracts/defi-fmint-minter.abi --pkg rpc --type DefiFMintMinter --out ./smc_fmint_minter.go

import (
	"github.com/ethereum/go-ethereum/common"
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
	contracts map[string]*common.Address
}

// tokenRegistryContract returns an instance of the fMint TokenRegistry
// smart contract used to keep track of tokens available on fMint protocol.
func (fmc *fMintConfig) tokenRegistryContract() (*DefiFMintTokenRegistry, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressTokenRegistry)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := NewDefiFMintTokenRegistry(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint TokenRegistry contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintMinterContract returns an instance of the fMint Minter smart contract.
func (fmc *fMintConfig) fMintMinterContract() (*DefiFMintMinter, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressMinter)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := NewDefiFMintMinter(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint Minter contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintRewardsDistribution returns an instance of the fMint Reward Distribution smart contract.
func (fmc *fMintConfig) fMintRewardsDistribution() (*FMintRewardsDistribution, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressRewardDistribution)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := NewFMintRewardsDistribution(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint Rewards Distribution contract; %s", err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintCollateralPool returns an instance of the fMint collateral pool contract.
func (fmc *fMintConfig) fMintTokenStorage(addr common.Address) (*DeFiTokenStorage, error) {
	// connect the contract
	contract, err := NewDeFiTokenStorage(addr, fmc.bridge.eth)
	if err != nil {
		fmc.bridge.log.Errorf("can not access fMint token pool %s; %s", addr.String(), err.Error())
		return nil, err
	}

	return contract, nil
}

// fMintCollateralPool returns an instance of the fMint collateral pool contract.
func (fmc *fMintConfig) fMintCollateralPool() (*DeFiTokenStorage, error) {
	// get address
	addr, err := fmc.contractAddress(fMintCollateralPool)
	if err != nil {
		return nil, err
	}

	return fmc.fMintTokenStorage(addr)
}

// fMintDebtPool returns an instance of the fMint debt pool contract.
func (fmc *fMintConfig) fMintDebtPool() (*DeFiTokenStorage, error) {
	// get address
	addr, err := fmc.contractAddress(fMintDebtPool)
	if err != nil {
		return nil, err
	}

	return fmc.fMintTokenStorage(addr)
}

// priceOracleProxyContract returns an instance of the DeFi price oracle proxy
// interface smart contract.
func (fmc *fMintConfig) priceOracleProxyContract() (*PriceOracleProxyInterface, error) {
	// get address
	addr, err := fmc.contractAddress(fMintAddressPriceOracleProxy)
	if err != nil {
		return nil, err
	}

	// connect the contract
	contract, err := NewPriceOracleProxyInterface(addr, fmc.bridge.eth)
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
	if nil == fmc.contracts {
		fmc.contracts = make(map[string]*common.Address, 6)
	}

	// lazy load the contract address requested
	if nil == fmc.contracts[name] {
		var err error
		fmc.contracts[name], err = fmc.loadAddress(name)
		if err != nil {
			return common.Address{}, err
		}
	}

	// return the address
	return *fmc.contracts[name], nil
}

// loadAddress loads a specified contract address from the AddressProvider.
func (fmc *fMintConfig) loadAddress(name string) (*common.Address, error) {
	// connect the Address Provider
	ap, err := NewDefiFMintAddressProvider(fmc.addressProvider, fmc.bridge.eth)
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
