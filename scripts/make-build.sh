#!/usr/bin/env bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

set -o errexit
set -o pipefail
set -o nounset
set -o xtrace

MODULE="github.com/kerria-dev/kerria"
COMMIT=$(git rev-parse --short HEAD)
DATE=$(date --iso-8601=seconds)

VERSION="dev"
CURRENT_REF=$(git rev-parse --abbrev-ref HEAD)
if [ "$CURRENT_REF" != "HEAD" ]; then
    VERSION="${VERSION} (${CURRENT_REF})"
fi
CURRENT_TAG=$(git describe --tags --exact-match || true)
if [[ -n "$CURRENT_TAG" ]]; then
    VERSION=$CURRENT_TAG
fi

${GO} build \
  -ldflags="
    -X '${MODULE}/cmd.version=${VERSION}'
    -X '${MODULE}/cmd.commit=${COMMIT}'
    -X '${MODULE}/cmd.date=${DATE}'" \
  -o="${GOBIN}/kerria"

