#!/bin/sh

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
