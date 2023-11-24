FROM 117503445/dev-base
LABEL maintainer="117503445"
# RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
WORKDIR /workspace

RUN pacman -Syu xorg-server-xvfb x11vnc --noconfirm

ENV SCREEN_WIDTH 1360
ENV SCREEN_HEIGHT 1020
ENV SCREEN_DEPTH 24
ENV DISPLAY :99.0

RUN git clone https://github.com/novnc/noVNC.git

EXPOSE 6080

RUN mkdir -p /workspace/container

RUN pacman -Syu python --noconfirm

RUN pacman -Syu plasma --noconfirm

ENTRYPOINT [ "./container/entrypoint.sh" ]

# RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist