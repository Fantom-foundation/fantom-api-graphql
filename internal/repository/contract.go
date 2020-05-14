/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"encoding/json"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"strings"
)

// contractCompilerInfoDelimiter is the delimiter used to separate
// different compiler information inside the contract detail
const contractCompilerInfoDelimiter = "; "

// Contract extract a smart contract information by account address, if available.
func (p *proxy) Contract(addr *common.Address) (*types.Contract, error) {
	return p.db.Contract(addr)
}

// Contracts returns list of smart contracts at Opera blockchain.
func (p *proxy) Contracts(validatedOnly bool, cursor *string, count int32) (*types.ContractList, error) {
	// go to the database for the list of contracts searched
	return p.db.Contracts(validatedOnly, cursor, count)
}

// compareContractCode compares provided compiled code with the transaction input.
func compareContractCode(tx *types.Transaction, code *string) (bool, error) {
	// decode the detail into byte array
	bc, err := hexutil.Decode(*code)
	if err != nil {
		return false, err
	}

	// is the transaction input shorter than the compiled contract?
	// if so there is no chance for pass
	if len(tx.InputData) < len(bc) {
		return false, nil
	}

	// get the length we need to compare
	n := len(bc)
	res := bytes.Compare(bc, tx.InputData[:n])

	// log the comparison entries and process result
	log.Printf("[:%d] %d = [%s]\n[%s]", n, res, hexutil.Bytes(bc).String(), tx.InputData[:n].String())

	// return the comparison result
	return res == 0, nil
}

// updateContractDetails updates local contract details from the provided compiler
// output.
func updateContractDetails(sc *types.Contract, detail *compiler.Contract) {
	// copy compiler information
	var str strings.Builder
	str.WriteString(detail.Info.Language)
	str.WriteString(" ")
	str.WriteString(detail.Info.LanguageVersion)
	str.WriteString(contractCompilerInfoDelimiter)
	str.WriteString(detail.Info.CompilerVersion)
	sc.Compiler = str.String()

	// copy ABI
	abi, err := json.Marshal(detail.Info.AbiDefinition)
	if err == nil {
		sc.Abi = string(abi)
	}

	// copy the source code
	sc.SourceCode = detail.Info.Source
}

// ValidateContract tries to validate contract byte code using
// provided source code. If successful, the contract information
// is updated the the repository.
func (p *proxy) ValidateContract(sc *types.Contract) error {
	// get the byte code of the actual contract
	tx, err := p.Transaction(&sc.TransactionHash)
	if err != nil {
		p.log.Errorf("can not get contract deployment transaction; %s", err.Error())
		return err
	}

	// try to compile the source code provided
	contracts, err := compiler.CompileSolidityString(p.solCompiler, sc.SourceCode)
	if err != nil {
		p.log.Errorf("solidity code compilation failed")
		return err
	}

	// loop over contracts ad try to validate one of them
	for name, detail := range contracts {
		// log the content
		p.log.Debugf("Sol %s: [%s]", name, detail.Code)

		// check if the compiled byte code match with the deployed contract
		match, err := compareContractCode(tx, &detail.Code)
		if err != nil {
			p.log.Errorf("contract byte code comparison failed")
			return err
		}

		// we have the winner
		if match {
			// set the contract name if not done already
			if 0 == len(sc.Name) {
				sc.Name = name
			}

			// update the contract data
			updateContractDetails(sc, detail)

			// write update to the database
			if err := p.db.UpdateContract(sc); err != nil {
				p.log.Errorf("contract validation failed due to db error; %s", err.Error())
				return err
			}

			// inform about success
			p.log.Debugf("contract %s validated with source %s", sc.Address.String(), name)

			// inform the upper instance we have a winner
			return nil
		}
	}

	// validation fails
	return fmt.Errorf("contract source code doesnt match with the deployed byte code")
}
