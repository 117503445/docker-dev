FROM archlinux
LABEL maintainer="117503445"
WORKDIR /root
COPY ./public ./public
RUN ./public/arch_init.sh
RUN ./public/download_vscode_server.sh
# RUN ./public/set_zsh.sh -p https://github.com/zsh-users/zsh-autosuggestions -p https://github.com/zsh-users/zsh-completions -p https://github.com/zsh-users/zsh-syntax-highlighting
# RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
WORKDIR /root/project
ENTRYPOINT [ "zsh" ]