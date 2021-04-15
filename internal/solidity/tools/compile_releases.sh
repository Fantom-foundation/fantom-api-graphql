#!/bin/bash
#
# This converts list of Solidity release version into a Go file used to control
# Solidity compiler access from the API.
#
# Please note, some runtime environments don't have access to all Solidity releases
# and can not use them to process incoming analyze requests sufficiently.
#
PACKAGE="solidity"
SOL_ORIGIN="https://github.com/ethereum/solidity.git"
OUT_FILE="$(dirname "$0")/../releases.go"

# check if the Solidity repository path was given
[ "$1" == "" ] && \
  echo "ERROR: Please specify Solidity GIT path." && \
  exit 1

# check if the given Solidity path exists
[ ! -d "$1" ] && \
  echo "ERROR: Expected Solidity repository not found on path ${1}." && \
  exit 1

# check if the given path is recognized as the expected GitHub Solidity repository
[ "$(git -C "${1}" config --get remote.origin.url)" != "${SOL_ORIGIN}" ] && \
  echo "ERROR: Expected Solidity repository not recognized on path ${1}." && \
  exit 1

# loop the revision tags
{
  # start the file and open the variable def
  echo "// Package solidity implements Solidity processor used to analyze"
  echo "// and verify Solidity based contracts."
  echo "package $PACKAGE"
  echo ""
  echo "// Auto generated Solidity releases list"
  echo "var solidityReleases = [...]string{"

  # list those version from the GIT repo
  for x in $(git -C "$1" tag -l --sort=-version:refname "v[01].[4-9].*"); do echo "    \"$x\","; done

  # close the variable
  echo "}"
} > "$OUT_FILE"
