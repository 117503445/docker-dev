FROM archlinux
LABEL maintainer="117503445"
WORKDIR /root
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syyu which zsh btop git vim wget gcc --noconfirm
COPY ./public/pacman.conf /etc/pacman.conf
RUN pacman-key --init && pacman -Syu archlinuxcn-keyring --noconfirm && pacman -S base-devel --noconfirm && pacman -S yay --noconfirm
COPY ./public ./public
RUN ./public/download_vscode_server.sh
RUN ./public/set_zsh.sh -p https://github.com/zsh-users/zsh-autosuggestions -p https://github.com/zsh-users/zsh-completions -p https://github.com/zsh-users/zsh-syntax-highlighting
COPY ./public/.p10k.zsh /root/.p10k.zsh
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
WORKDIR /root/project
ENTRYPOINT [ "zsh" ]