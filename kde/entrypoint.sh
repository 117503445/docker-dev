#!/usr/bin/env bash

# Start Xvfb
export GEOMETRY="$SCREEN_WIDTH""x""$SCREEN_HEIGHT""x""$SCREEN_DEPTH"
Xvfb $DISPLAY -screen 0 $GEOMETRY -ac +extension RANDR > /workspace/container/xvfb.stdout.log 2>&1 &

# Start fluxbox
pacman -Syu fluxbox --noconfirm
fluxbox -display $DISPLAY > /workspace/container/fluxbox.stdout.log 2>&1 &

# start x11
x11vnc -display $DISPLAY -nopw -listen localhost -xkb -ncache 10 -ncache_cr -forever > /workspace/container/x11vnc.stdout.log 2>&1 &

# Start noVNC
/workspace/noVNC/utils/novnc_proxy --vnc localhost:5900 --listen 0.0.0.0:6080 > /workspace/container/noVNC.stdout.log 2>&1 &

# startplasma-x11 > /workspace/container/plasma.stdout.log 2>&1 &



zsh