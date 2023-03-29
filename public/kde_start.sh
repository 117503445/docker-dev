#!/bin/bash

set -exv

# current user
CUSER=${KDE_USERNAME:-root}
# add user if specified
if [ ! -z $KDE_USERNAME ]; then
	HOMEDIR="/home/$KDE_USERNAME"
	# Check if user exists
	if ! id "$KDE_USERNAME" &> /dev/null; then
		pacman -S --noconfirm sudo
		useradd -m -G wheel $KDE_USERNAME
		# delete password
		passwd -d $KDE_USERNAME
		echo "%wheel ALL=(ALL:ALL) ALL" >> /etc/sudoers
		su $KDE_USERNAME
	fi
else
	HOMEDIR="/root"
	passwd -d root
fi

umask 0077                # use safe default permissions
mkdir -p "$HOMEDIR/.vnc"
chmod go-rwx "$HOMEDIR/.vnc" # enforce safe permissions

# chromium
sed -i 's/\/usr\/bin\/chromium/\/usr\/bin\/chromium --no-sandbox/g' /usr/share/applications/chromium.desktop

# Start TigerVNC
if [ ! -z $VNC_PASSWD ]; then
	vncpasswd -f <<< "$VNC_PASSWD" > "$HOMEDIR/.vnc/passwd"
fi

# chown -R $CUSER:$CUSER "$HOMEDIR"
# Remove lock since stopping containers won't remove it
rm -f /tmp/.X0-lock

echo Starting vncsession...
vncsession $CUSER :0

# start kde
unset SESSION_MANAGER && unset DBUS_SESSION_BUS_ADDRESS && DISPLAY=:0 exec dbus-launch startplasma-x11 > "$HOME/kde.log" 2>&1 &

# Start noVNC
if [ "$DISABLE_NOVNC" != "true" ]; then
	/noVNC/utils/launch.sh
else
	# prevent process from exiting
	tail -f /dev/null
fi