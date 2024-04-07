FROM 117503445/dev-base
LABEL maintainer="117503445"
ENV PATH="/root/go/bin:${PATH}"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu go --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

WORKDIR /scripts
COPY ./scripts/set_go.sh .
RUN ./set_go.sh && rm set_go.sh
WORKDIR /workspace

ENV VSC_EXTS="golang.go,${VSC_EXTS}"