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
"${CMD_LS_FILES[@]}" -z "pkg/**/${GEN_FILE_BASE}.*.go" \
  | xargs --null rm --force
