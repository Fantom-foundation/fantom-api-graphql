package repository

import "github.com/ethereum/go-ethereum/common"

// Erc165SupportsInterface provides information about support of interface in the contract.
func (p *proxy) Erc165SupportsInterface(address *common.Address, interfaceID [4]byte) (bool, error) {
	return p.rpc.Erc165SupportsInterface(address, interfaceID)
}
