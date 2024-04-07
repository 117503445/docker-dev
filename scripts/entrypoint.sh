#!/usr/bin/env bash

set -e

curl -s https://vsc-server.unidrop.top/vsc_server_setup.sh | sh

pacman -Syu --noconfirm

# check if stdin is a terminal
if [ -t 0 ]; then
  exec bash
else
  exec tail -f /dev/null
fi