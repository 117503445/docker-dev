FROM archlinux
LABEL maintainer="117503445"
WORKDIR /scripts
COPY ./scripts/arch_init.sh .
RUN ./arch_init.sh && rm arch_init.sh

COPY ./scripts/download_vscode_server.sh .
RUN ./download_vscode_server.sh && rm download_vscode_server.sh

COPY ./scripts/install_vsc_ext.py .
RUN ./install_vsc_ext.py mhutchie.git-graph ms-ceintl.vscode-language-pack-zh-hans github.copilot christian-kohler.path-intellisense

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

WORKDIR /workspace
ENTRYPOINT [ "fish" ]