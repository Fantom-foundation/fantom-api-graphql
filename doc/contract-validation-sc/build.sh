#!/bin/bash
#
# Solidity compiler must be installed
#
rm -f ./*.bin ./*.abi
solc -o ./ --optimize --optimize-runs=200 --abi --bin ./ContractValidation.sol
