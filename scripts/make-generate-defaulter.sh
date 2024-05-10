#!/usr/bin/env bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

# Setup variables
source ./scripts/make-vars.sh

set -o errexit
set -o pipefail
set -o nounset
set -o xtrace

# Delete previously generated files
"${CMD_LS_FILES[@]}" -z "pkg/**/${GEN_FILE_BASE_DEFAULTER}.go" \
  | xargs -0 rm -f

# Execute codegen
"${CMD_DEFAULTER_GEN[@]}" \
    "${GEN_COMMON_FLAGS[@]}" \
    --output-file-base "$GEN_FILE_BASE_DEFAULTER" \

# HACK: The generator tries to register each defaulter with the scheme.
#       We don't support this, since we don't generate deep copy or use it anyway.
TARGET_FILE="${GEN_API_V1ALPHA1}/${GEN_FILE_BASE_DEFAULTER}.go"
TARGET_FUNCTION="RegisterDefaults"
sed -i "/\/\/.*${TARGET_FUNCTION}/,/^}/d" "${TARGET_FILE}"
"${GOBIN}/goimports" -w "${TARGET_FILE}"
"$("${GO}" env GOROOT)/bin/gofmt" -w "${TARGET_FILE}"
