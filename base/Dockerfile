FROM archlinux
LABEL maintainer="117503445"
WORKDIR /scripts
COPY ./scripts/arch_init.sh .
RUN ./arch_init.sh && rm arch_init.sh

ENV VSC_EXTS mhutchie.git-graph,ms-ceintl.vscode-language-pack-zh-hans,github.copilot,christian-kohler.path-intellisense

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

WORKDIR /workspace

COPY ./scripts/entrypoint.py /entrypoint.py
COPY ./scripts/vsc_server_setup.py /vsc_server_setup.py

ENTRYPOINT [ "/entrypoint.py" ]