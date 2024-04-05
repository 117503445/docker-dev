#!/usr/bin/env bash

set -e

# check if stdin is a terminal
if [ -t 0 ]; then
  exec fish
else
  exec tail -f /dev/null
fi