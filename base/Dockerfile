FROM archlinux
LABEL maintainer="117503445"

# 根据是否存在 $CHINA_MIRROR 环境变量设置软件源
ARG CHINA_MIRROR
RUN if [ -z "$CHINA_MIRROR" ]; then \
      echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist; \
    else \
      echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist; \
    fi

# 初始化并填充 pacman 密钥环
RUN pacman-key --init && \
    pacman-key --populate && \
    pacman -Sy archlinux-keyring --noconfirm && \
    pacman -Su --noconfirm

# 安装所需软件包
RUN pacman -Syu which zsh fish btop git openssh docker docker-compose docker-buildx nano vim micro base-devel parted tmux python wget yazi go-task go --noconfirm

# 如果设置了 $CHINA_MIRROR 环境变量，则配置go代理
RUN if [ -n "$CHINA_MIRROR" ]; then \
      go env -w GOPROXY=https://goproxy.cn,direct; \
    fi

# 添加用户 builder，并安装 yay AUR 助手
RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay.git yay && cd yay && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

RUN su - builder -c "yay -Su code-server --noconfirm"
EXPOSE 4444

# 配置 git 全局设置
RUN git config --global user.name "117503445" && \
    git config --global user.email t117503445@gmail.com && \
    git config --global push.default current && \
    git config --global core.editor "code --wait"

RUN go install github.com/mattn/goreman@latest
ENV PATH="/root/go/bin:${PATH}"

RUN go env -w GOPROXY=direct
RUN go install github.com/117503445/vsc-init@master && vsc-init

# 更改默认 shell 为 fish
RUN chsh -s /usr/bin/fish
COPY ./scripts/config.fish /root/.config/fish/config.fish

RUN mkdir -p /docker-dev/logs
COPY ./entrypoint /tmp/entrypoint
RUN cd /tmp/entrypoint && cp Procfile /docker-dev/Procfile && go build . && cp ./entrypoint /docker-dev/entrypoint && rm -rf /tmp/entrypoint

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace

ENTRYPOINT [ "/docker-dev/entrypoint" ]