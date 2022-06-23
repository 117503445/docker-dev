FROM archlinux
LABEL maintainer="117503445"
WORKDIR /root
COPY ./public ./public
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
# RUN echo "Server = https://mirrors.aliyun.com/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syyu which zsh btop git vim wget gcc --noconfirm
RUN ./public/download_vscode_server.sh
WORKDIR /root/project
RUN echo "Server = https://mirrors.aliyun.com/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
ENTRYPOINT [ "zsh" ]