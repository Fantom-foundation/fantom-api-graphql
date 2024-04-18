#!/bin/bash
#
# This bundles the GraphQL schema tree into a single source file so it can be complied with the code base
#
PACKAGE="schema"
FILE="schema.graphql"
BASE_FOLDER="$(dirname "$0")/.."

# Make the bundle file content by combining the *.graphql files
{
  echo "## Auto generated GraphQL schema bundle"
  find "$BASE_FOLDER/definition" -name '*.graphql' -print0 | xargs -0 -I{} sh -c "cat {}; echo ''"
} >"$BASE_FOLDER/$FILE"
