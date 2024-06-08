FROM archlinux
LABEL maintainer="117503445"


WORKDIR /scripts
COPY ./scripts/arch_init.sh .
RUN ./arch_init.sh && rm arch_init.sh

RUN cd /tmp && git clone https://github.com/117503445/docker-dev.git --depth=1 && cd docker-dev/entrypoint && go build . && cp ./entrypoint /entrypoint

ENV VSC_EXTS mhutchie.git-graph,eamodio.gitlens,ms-ceintl.vscode-language-pack-zh-hans,github.copilot,christian-kohler.path-intellisense,tamasfe.even-better-toml,redhat.vscode-yaml

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

WORKDIR /workspace

# COPY ./scripts/entrypoint.py /entrypoint.py
COPY ./scripts/vsc_server_setup.py /vsc_server_setup.py

# HEALTHCHECK --interval=3s --timeout=1s --start-period=240s --retries=5 CMD [ "test", "!", "-f", "/entrypoint.py.lock" ]

ENTRYPOINT [ "/entrypoint" ]