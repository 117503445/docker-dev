#!/bin/bash
set -exv

umask 0077                # use safe default permissions
chmod go-rwx "$HOME/.vnc" # enforce safe permissions
chmod +x "$HOME/.vnc/xstartup"

# chromium
sed -i 's/\/usr\/bin\/chromium/\/usr\/bin\/chromium --no-sandbox/g' /usr/share/applications/chromium.desktop

# Start TigerVNC
# if [ ! -z $VNC_PASSWD ]; then
# 	vncpasswd -f <<< "$VNC_PASSWD" > "$HOME/.vnc/passwd"
# 	vncserver -rfbport 5900
# else
# 	vncpasswd -f <<< "" > "$HOME/.vnc/passwd"
# 	vncserver -rfbport 5900 -SecurityTypes None
# fi


# x11vnc -storepasswd password ~/.vnc/passwd

# $DISPLAY=0
# [ -x ~/.vnc/xstartup ] && DISPLAY=0 exec ~/.vnc/xstartup
# startx

# echo "exec dbus-launch startplasma-x11" > ~/.xinitrc && chmod +x ~/.xinitrc


# x11vnc -rfbauth ~/.vnc/passwd -display :0
# nohup x11vnc  -rfbport 5900 -rfbauth ~/.vnc/passwd -create -forever  > x11vnc.log &
# x11vnc  -rfbport 5900 -rfbauth ~/.vnc/passwd -create -forever


# vncpasswd -f <<< "password" > "$HOME/.vnc/passwd"
chmod 600 $HOME/.vnc/passwd
x11vnc -storepasswd password /root/.vnc/passwd
DISPLAY=:1 sh ~/.vnc/xstartup &
x11vnc  -rfbport 5900 -rfbauth /root/.vnc/passwd -forever &

# Start noVNC
/noVNC-${noVNC_version}/utils/launch.sh

 vncserver :0
 dbus-daemon --config-file=/usr/share/dbus-1/system.conf