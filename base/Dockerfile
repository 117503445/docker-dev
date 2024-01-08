FROM archlinux
LABEL maintainer="117503445"
WORKDIR /root
COPY ./public/arch_init.sh ./public/arch_init.sh 
RUN ./public/arch_init.sh

COPY ./public/download_vscode_server.sh ./public/download_vscode_server.sh
RUN ./public/download_vscode_server.sh

COPY ./public/install_vsc_ext.py ./public/install_vsc_ext.py
RUN ./public/install_vsc_ext.py mhutchie.git-graph ms-ceintl.vscode-language-pack-zh-hans github.copilot christian-kohler.path-intellisense

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

COPY ./public ./public
WORKDIR /workspace
ENTRYPOINT [ "fish" ]