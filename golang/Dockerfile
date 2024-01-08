FROM 117503445/dev-base
LABEL maintainer="117503445"
ENV PATH="/root/go/bin:${PATH}"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu go --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN /root/public/set_go.sh
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN ./public/install_vsc_ext.py golang.go