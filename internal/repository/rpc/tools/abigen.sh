#!/bin/bash
#
# Generate the contract ABI code using the right ABI generator tool from bin dir
#
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
