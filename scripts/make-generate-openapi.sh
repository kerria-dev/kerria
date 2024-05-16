#!/usr/bin/env bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

# Setup variables
source ./scripts/make-vars.sh

set -o errexit
set -o pipefail
set -o nounset
set -o xtrace

delete_generated_files() {
  "${CMD_LS_FILES[@]}" -z "pkg/**/${GEN_FILE_BASE_OPENAPI}.go" \
    | xargs --null rm --force
}

# Delete previously generated files
delete_generated_files

# Execute codegen
ALLOWED_VIOLATIONS=(
  '^API rule violation: [A-Za-z_]+,k8s\.io/apimachinery.*$'
)
VIOLATION_PATTERN=$(IFS="|"; echo "${ALLOWED_VIOLATIONS[*]}")
if "${CMD_OPENAPI_GEN[@]}" \
    "${GEN_COMMON_FLAGS[@]}" \
    --input-dirs "${GEN_GO_MODULE}/${GEN_API_META}" \
    --output-file-base "$GEN_FILE_BASE_OPENAPI" \
    --output-package "${GEN_GO_MODULE}/${GEN_OPENAPI}" \
  | grep --invert-match --extended-regexp "$VIOLATION_PATTERN"; then
    delete_generated_files
    exit 1;
fi
