#!/usr/bin/env bash

set -e

echo "=== Install VSC Server & Extensions ==="
curl -s https://vsc-server.unidrop.top/vsc_server_setup.sh | sh
echo "=== Install VSC Server & Extensions Done ==="

echo "=== Update Arch Linux Packages ==="
pacman -Syu --noconfirm
echo "=== Update Arch Linux Packages Done ==="

echo "~~~ Successfully Init Development Environment, Enjoy! ~~~"

# check if stdin is a terminal
if [ -t 0 ]; then
  exec bash
else
  exec tail -f /dev/null
fi