FROM archlinux
LABEL maintainer="117503445"
WORKDIR /root
COPY ./public/arch_init.sh ./public/arch_init.sh 
RUN ./public/arch_init.sh

COPY ./public/download_vscode_server.sh ./public/download_vscode_server.sh
RUN ./public/download_vscode_server.sh
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

# COPY ./public/vsc_init.sh ./public/vsc_init.sh
# RUN export COMMIT_SHA=$(./public/vsc_init.sh) && export PATH=~/.vscode-server/bin/${COMMIT_SHA}/bin/remote-cli:$PATH
# RUN ./public/vsc_init.sh

COPY ./public ./public
WORKDIR /workspace
ENTRYPOINT [ "zsh" ]