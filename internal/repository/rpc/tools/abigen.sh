#!/bin/bash
#
# Generate the contract ABI code using the right ABI generator tool from bin dir
#
# You need to install abigen to use this tool:
# 1) checkout go-ethereum repository
# 2) run "go install ./cmd/abigen" (or "make devtools")
# 3) the abigen tool should be now installed in $GOPATH
#
if [ -z "${GOPATH}" ]; then GOPATH="${HOME}/go"; fi
ABIGEN="${GOPATH}/bin/abigen"

if [ ! -x "${ABIGEN}" ]
then
  echo "ABI generator not found as [${ABIGEN}]."
  exit 1
fi

"${ABIGEN}" "${@}"
exit $?
