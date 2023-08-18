FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu rustup --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
COPY rust/cargo_config /root/.cargo/config
RUN rustup default stable
