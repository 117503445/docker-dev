#!/bin/bash

set -exv

echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

# Install apps
# pacman-key --init && pacman -Sy archlinux-keyring --noconfirm &&
pacman -Syu --noconfirm plasma-meta \
	kde-accessibility-meta kde-system-meta konsole \
	chromium vim wget tigervnc xorg-server \
	python-numpy python-setuptools git \
	&& pacman -Scc --noconfirm

# Install noVNC
if [ "$DISABLE_NOVNC" != "true" ]; then
	export noVNC_version=1.2.0
	export websockify_version=0.10.0

	wget https://github.com/novnc/websockify/archive/v${websockify_version}.tar.gz -O /websockify.tar.gz \
		&& tar -xvf /websockify.tar.gz -C / \
		&& cd /websockify-${websockify_version} \
		&& python setup.py install \
		&& cd / && rm -r /websockify.tar.gz /websockify-${websockify_version} \
		&& wget https://github.com/novnc/noVNC/archive/v${noVNC_version}.tar.gz -O /noVNC.tar.gz \
		&& tar -xvf /noVNC.tar.gz -C / \
		&& mv /noVNC-${noVNC_version} /noVNC \
		&& cd /noVNC \
		&& ln -s vnc.html index.html \
		&& rm /noVNC.tar.gz
fi
pacman -Syu kwallet-pam noto-fonts-cjk --noconfirm

echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist