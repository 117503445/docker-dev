ARG BASE_IMAGE=archlinux

FROM $BASE_IMAGE
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
RUN pacman -Syu which zsh fish btop git openssh docker docker-compose docker-buildx nano vim micro base-devel parted tmux python wget yazi go-task go zsh-autosuggestions zsh-syntax-highlighting podman skopeo umoci ncdu cloc --noconfirm

# 添加用户 builder，并安装 yay AUR 助手
RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay-bin.git && cd yay-bin && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

ENV PATH="/root/go/bin:${PATH}"
RUN go env -w GOPROXY=direct
RUN go install github.com/117503445/vsc-init@master
RUN su - builder -c "yay -Su code-server --noconfirm"
EXPOSE 4444
RUN /root/go/bin/vsc-init

# 如果设置了 $CHINA_MIRROR 环境变量，则配置go代理
RUN if [ -n "$CHINA_MIRROR" ]; then \
      go env -w GOPROXY=https://goproxy.cn,direct; \
    else \
      go env -w GOPROXY=https://proxy.golang.org,direct; \
    fi

# 配置 git 全局设置
RUN git config --global user.name "117503445" && \
    git config --global user.email t117503445@gmail.com && \
    git config --global push.default current && \
    git config --global core.editor "code --wait"

RUN go install github.com/mattn/goreman@latest
RUN go install github.com/ycd/dstp/cmd/dstp@latest

# 更改默认 shell 为 zsh
RUN chsh -s /usr/bin/zsh
# RUN su - builder -c "yay -Su --noconfirm oh-my-zsh-git"
RUN if [ -z "$CHINA_MIRROR" ]; then \
      sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"; \
    else \
      sh -c "$(curl -fsSL https://install.ohmyz.sh)"; \
    fi
COPY ./scripts/.zshrc /root/.zshrc

RUN go install github.com/cweill/gotests/gotests@latest && go install github.com/fatih/gomodifytags@latest && go install github.com/josharian/impl@latest && go install github.com/haya14busa/goplay/cmd/goplay@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install honnef.co/go/tools/cmd/staticcheck@latest && go install golang.org/x/tools/gopls@latest && go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/goreleaser/goreleaser/v2@latest

# front
RUN pacman -Syu npm yarn pnpm --noconfirm
RUN npm config set registry https://registry.npmmirror.com
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn global add typescript
RUN pnpm config set registry https://registry.npmmirror.com

# python
RUN pacman -Syu poetry mypy rye uv --noconfirm
RUN poetry config virtualenvs.in-project true
RUN su - builder -c "yay -Su miniconda3 --noconfirm"
COPY ./dev/python/.condarc /root/.condarc
COPY ./dev/python/uv.toml /root/.config/uv/uv.toml

# latex
# RUN pacman -Syu texlive make gnu-free-fonts texlive-langchinese biber --noconfirm
# RUN git clone https://github.com/Haixing-Hu/latex-chinese-fonts.git /usr/share/fonts/latex-chinese-fonts --depth=1

# cpp
RUN pacman -Syu gcc clang cmake --noconfirm

# csharp
RUN pacman -Syu --noconfirm dotnet-runtime dotnet-sdk aspnet-runtime

# java
RUN pacman -Sy jdk-openjdk maven gradle kotlin --noconfirm
# RUN su - builder -c "yay -Su scala --noconfirm"

# rust
RUN pacman -Syu rust-analyzer rustup --noconfirm
COPY ./dev/rust/config.toml /root/.cargo/config.toml
RUN sh -c "if [ -n \"$CHINA_MIRROR\" ]; then \
                export RUSTUP_DIST_SERVER=https://rsproxy.cn; \
                export RUSTUP_UPDATE_ROOT=https://rsproxy.cn/rustup; \
                rustup default stable; \
            else \
                rustup default stable; \
            fi"
RUN cargo install cargo-watch
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
ENV PATH="/root/.cargo/bin:$PATH"

# typst
# RUN pacman -Syu typst noto-fonts noto-fonts-cjk adobe-source-han-sans-cn-fonts noto-fonts-emoji ttf-ubuntu-font-family --noconfirm
# COPY typst/fonts /usr/share/fonts/docker-dev
# ENV TYPST_FONT_PATHS=/usr/share/fonts
# RUN wget https://github.com/Enter-tainer/typstyle/releases/latest/download/typstyle-x86_64-unknown-linux-gnu -o typstyle && chmod +x typstyle && mv typstyle /usr/bin/typstyle

RUN mkdir -p /docker-dev/logs
COPY ./entrypoint /tmp/entrypoint
RUN cd /tmp/entrypoint && cp Procfile /docker-dev/Procfile && go build . && cp ./entrypoint /docker-dev/entrypoint && rm -rf /tmp/entrypoint

WORKDIR /workspace

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN curl https://wiki.117503445.top/linux/script/ssh.sh | bash
RUN ssh-keygen -A

ENV TZ=Asia/Shanghai

ENTRYPOINT [ "/docker-dev/entrypoint" ]