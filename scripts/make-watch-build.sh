#!/bin/sh

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

while true;
do
  if ! make bin/kerria -q "$@";
  then
    echo "#-> Starting build: $(date)"
    make bin/kerria "$@";
    echo "#-> Build complete."
  fi
  sleep 0.5;
done
