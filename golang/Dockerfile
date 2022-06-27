FROM 117503445/dev-base
LABEL maintainer="117503445"
ENV PATH="/root/go/bin:${PATH}"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -S go --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN /root/public/set_go.sh
RUN wget -O gf -q https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf