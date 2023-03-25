FROM 117503445/dev-base
LABEL maintainer="117503445"
ENV PATH="/root/go/bin:${PATH}"
# RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

ENV noVNC_version=1.2.0
ENV websockify_version=0.9.0

# Install apps
# RUN pacman -Syu --noconfirm plasma-meta \
# 	kde-accessibility-meta kde-system-meta konsole \
# 	chromium vim wget tigervnc xorg-server \
# 	python-numpy python-setuptools \
# 	&& pacman -Scc --noconfirm
RUN pacman -Syu --noconfirm plasma-meta \
	kde-accessibility-meta kde-system-meta konsole \
	chromium vim wget x11vnc tigervnc xorg-server \
	python-numpy python-setuptools \
	&& pacman -Scc --noconfirm
# Install noVNC
RUN	wget https://github.com/novnc/websockify/archive/v${websockify_version}.tar.gz -O /websockify.tar.gz \
	&& tar -xvf /websockify.tar.gz -C / \
	&& cd /websockify-${websockify_version} \
	&& python setup.py install \
	&& cd / && rm -r /websockify.tar.gz /websockify-${websockify_version} \
	&& wget https://github.com/novnc/noVNC/archive/v${noVNC_version}.tar.gz -O /noVNC.tar.gz \
	&& tar -xvf /noVNC.tar.gz -C / \
	&& cd /noVNC-${noVNC_version} \
	&& ln -s vnc.html index.html \
	&& rm /noVNC.tar.gz

COPY ./public/xstartup /root/.vnc/
# COPY ./start.sh /

WORKDIR /root

EXPOSE 5900 6080

CMD [ "/root/public/kde-start.sh" ]