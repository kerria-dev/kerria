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
"${CMD_LS_FILES[@]}" -z "pkg/**/${GEN_FILE_BASE_OPENAPI}.go" \
  | xargs --null rm --force

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
  exit 1;
fi
