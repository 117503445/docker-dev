FROM archlinux

# COPY ./test/entrypoint.sh /entrypoint.sh

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

RUN pacman -Syu fish docker docker-compose docker-buildx fuse-overlayfs --noconfirm

COPY ./test/entrypoint.sh /docker-dev/entrypoint

WORKDIR /workspace

# ENTRYPOINT [ "/docker-dev/entrypoint" ]

# ENTRYPOINT [ "tail", "-f", "/dev/null" ]
ENTRYPOINT [ "dockerd" ]