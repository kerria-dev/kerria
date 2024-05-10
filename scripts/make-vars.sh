#!/usr/bin/env bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

export CMD_LS_FILES=(git ls-files --cached --modified --others --exclude-standard)
export CMD_DEFAULTER_GEN=("$GOBIN"/defaulter-gen)
export CMD_OPENAPI_GEN=("$GOBIN"/openapi-gen)

export GEN_GO_MODULE='github.com/kerria-dev/kerria'
export GEN_APIS='pkg/apis'
export GEN_API_META="${GEN_APIS}/kerria.dev/meta"
export GEN_API_V1ALPHA1="${GEN_APIS}/kerria.dev/v1alpha1"
export GEN_OPENAPI='pkg/openapi'
export GEN_FILE_BASE='zz_generated'
export GEN_FILE_BASE_DEFAULTER="${GEN_FILE_BASE}.defaults"
export GEN_FILE_BASE_OPENAPI="${GEN_FILE_BASE}.openapi"
export GEN_COMMON_FLAGS=(
    --input-dirs "${GEN_GO_MODULE}/${GEN_API_V1ALPHA1}"
    --output-base '.'
    --go-header-file 'hack/boilerplate.go.txt'
    --trim-path-prefix "${GEN_GO_MODULE}"
    --v 2
)

GEN_API_SOURCES=$(
    # shellcheck disable=SC2038
    find "${GEN_APIS}" -type f -not -name "${GEN_FILE_BASE}*" \
      | xargs -I {} echo -n "{} "
  )
export GEN_API_SOURCES

if [ -v SOURCE_DIRS ]; then
  IFS=' ' read -r -a SOURCE_DIRS_ARRAY <<< "${SOURCE_DIRS}"
  BUILD_SOURCES=$(
      # shellcheck disable=SC2038
      find "${SOURCE_DIRS_ARRAY[@]}" -type f -not -name "${GEN_FILE_BASE}*" \
        | xargs -I {} echo -n "{} "
    )
  export BUILD_SOURCES
fi
