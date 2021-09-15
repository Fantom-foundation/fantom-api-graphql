#!/bin/bash
#
# Generate the contract ABI code using the right ABI generator tool from bin dir
#
# You need to install abigen to use this tool:
# 1) checkout go-ethereum repository
# 2) run "go install ./cmd/abigen" (or "make devtools")
# 3) the abigen tool should be now installed in $GOPATH
#

if [ -x "${GOPATH}"/bin/abigen ]
then
  "${GOPATH}"/bin/abigen "${@}"
  exit $?
fi

BASE_FOLDER=$( cd "$(dirname "$0")"/../../../.. && pwd)
OS=$(uname)
case $OS in
'Linux')
  "${BASE_FOLDER}"/tools/bin/abigen-linux-amd64 "${@}"
  exit $?
  ;;
'Darwin')
  "${BASE_FOLDER}"/tools/bin/abigen-darwin-arm64 "${@}"
  exit $?
  ;;
*)
  echo "Can not determine the right ABI generator tool."
  exit 1
  ;;
esac
