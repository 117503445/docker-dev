# ============================================
# Build Stages
# ============================================
ARG GO_BASE_IMAGE=golang:1.24.1
ARG BASE_IMAGE=lscr.io/linuxserver/webtop:arch-kde

# Build entrypoint
FROM ${GO_BASE_IMAGE} AS builder-entrypoint
WORKDIR /build
COPY scripts/entrypoint/go.mod scripts/entrypoint/go.sum ./
RUN go mod download
COPY scripts/entrypoint/ ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

# Build vsc-init
FROM ${GO_BASE_IMAGE} AS builder-vsc-init
WORKDIR /build
COPY scripts/vsc-init/go.mod scripts/vsc-init/go.sum ./
RUN go mod download
COPY scripts/vsc-init/ ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

# Build shell-init
FROM ${GO_BASE_IMAGE} AS builder-shell-init
WORKDIR /build
COPY scripts/shell-init/go.mod scripts/shell-init/go.sum ./
RUN go mod download
COPY scripts/shell-init/ ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

# Build vibe-init
FROM ${GO_BASE_IMAGE} AS builder-vibe-init
WORKDIR /build
COPY scripts/vibe-init/go.mod scripts/vibe-init/go.sum ./
RUN go mod download
COPY scripts/vibe-init/ ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app .

# ============================================
# Main Stage
# ============================================

FROM $BASE_IMAGE
LABEL maintainer="117503445"

USER root
ENV PUID=0
ENV GUID=0
ENV HOME=/root

RUN locale-gen en_US.UTF-8
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

ENV TZ=Asia/Shanghai

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
RUN pacman -Syu which zsh fish btop git openssh docker docker-compose docker-buildx nano vim micro base-devel parted tmux python wget yazi go-task go zsh-autosuggestions zsh-syntax-highlighting podman skopeo umoci ncdu cloc fuse-overlayfs nvidia-container-toolkit github-cli jq rsync kubectl helm unzip npm --noconfirm

# 添加用户 builder，并安装 yay AUR 助手
RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay-bin.git && cd yay-bin && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

# Install AUR packages
RUN su - builder -c "yay -Su code-server claude-code cursor-cli openai-codex-bin --noconfirm"

# latex
RUN git clone https://github.com/Haixing-Hu/latex-chinese-fonts.git /usr/share/fonts/latex-chinese-fonts --depth=1

# cpp
RUN pacman -Syu gcc clang cmake --noconfirm

# csharp
RUN pacman -Syu --noconfirm dotnet-runtime dotnet-sdk aspnet-runtime

# java
RUN pacman -Sy jdk-openjdk maven gradle kotlin --noconfirm

# typst
RUN pacman -Syu typst noto-fonts noto-fonts-cjk adobe-source-han-sans-cn-fonts noto-fonts-emoji ttf-ubuntu-font-family --noconfirm
ENV TYPST_FONT_PATHS=/usr/share/fonts
RUN wget https://github.com/Enter-tainer/typstyle/releases/latest/download/typstyle-x86_64-unknown-linux-gnu -O typstyle && chmod +x typstyle && mv typstyle /usr/bin/typstyle || echo "Warning: Failed to download typstyle"

# aliyun
RUN npm install @serverless-devs/s -g
RUN wget https://gosspublic.alicdn.com/ossutil/1.7.19/ossutil-v1.7.19-linux-amd64.zip && unzip ossutil-v1.7.19-linux-amd64.zip && mv ossutil-v1.7.19-linux-amd64/ossutil /usr/local/bin/ && rm ossutil-v1.7.19-linux-amd64.zip ossutil-v1.7.19-linux-amd64/ossutil64 && rm -r ossutil-v1.7.19-linux-amd64 || echo "Warning: Failed to download ossutil"

# python
RUN pacman -Syu uv --noconfirm
COPY ./assets/python/uv.toml /tmp/uv.toml
RUN mkdir -p $HOME/.config/uv && mv /tmp/uv.toml $HOME/.config/uv/uv.toml

# front
RUN pacman -Syu npm yarn pnpm --noconfirm
RUN npm config set registry https://registry.npmmirror.com
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn global add typescript
RUN pnpm config set registry https://registry.npmmirror.com

# rust
RUN pacman -Syu rust-analyzer --noconfirm
RUN cargo install cargo-watch
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
ENV PATH="$HOME/.cargo/bin:${PATH}"
COPY ./assets/rust/config.toml /tmp/cargo-config.toml
RUN mkdir -p $HOME/.cargo && mv /tmp/cargo-config.toml $HOME/.cargo/config.toml

# Setup Go environment
ENV PATH="$HOME/go/bin:${PATH}"
RUN go env -w GOPATH=$HOME/go

EXPOSE 4444

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
RUN go install github.com/117503445/histmon@master

RUN go install github.com/cweill/gotests/gotests@latest && go install github.com/fatih/gomodifytags@latest && go install github.com/josharian/impl@latest && go install github.com/haya14busa/goplay/cmd/goplay@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install honnef.co/go/tools/cmd/staticcheck@latest && go install golang.org/x/tools/gopls@latest && go install github.com/ramya-rao-a/go-outline@latest && go install github.com/goreleaser/goreleaser/v2@latest && go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest && go install -v github.com/incu6us/goimports-reviser/v3@latest && go install github.com/bufbuild/buf/cmd/buf@latest && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

RUN mkdir -p /docker-dev/logs

WORKDIR /workspace

RUN curl https://wiki.117503445.top/linux/script/ssh.sh | bash
COPY ./assets/sshd/ssh_host_ed25519_key /etc/ssh/ssh_host_ed25519_key
COPY ./assets/sshd/ssh_host_ed25519_key.pub /etc/ssh/ssh_host_ed25519_key.pub
RUN chmod 600 /etc/ssh/ssh_host_ed25519_key

ENV TZ=Asia/Shanghai
ENV DOCKER_BUILDKIT=1
ENV PATH="$HOME/.local/bin:${PATH}"

# 配置 custom service 启动 /docker-dev/entrypoint (linuxserver 镜像推荐方式)
RUN mkdir -p /custom-services.d && \
    printf '#!/usr/bin/with-contenv bash\nexec /docker-dev/entrypoint\n' > /custom-services.d/docker-dev-entrypoint && \
    chmod +x /custom-services.d/docker-dev-entrypoint

# 创建 KDE 桌面模板文件，使右键菜单可以创建新文件
RUN mkdir -p $HOME/Templates && \
    touch "$HOME/Templates/Text File.txt" && \
    touch "$HOME/Templates/Empty File"

# 配置 abc 用户以 root 权限运行桌面，并设置桌面路径为 /root/Desktop
RUN usermod -d /root abc && \
    mkdir -p /root/Desktop && \
    chown -R abc:abc /root && \
    chmod 755 /root /root/Desktop

# 配置 KDE Plasma 桌面 Folder View 显示 /root/Desktop 内容
RUN mkdir -p /root/.config/autostart /usr/local/bin && \
    printf '#!/bin/bash\n\
CONFIG="/root/.config/plasma-org.kde.plasma.desktop-appletsrc"\n\
sleep 3\n\
\n\
if [ -f "$CONFIG" ] && ! grep -q "url=/root/Desktop" "$CONFIG"; then\n\
  if ! grep -q "plugin=org.kde.plasma.folder" "$CONFIG"; then\n\
    sed -i "/^\\[Containments\\]\\[1\\]$/a plugin=org.kde.plasma.folder" "$CONFIG"\n\
  fi\n\
  \n\
  if ! grep -q "\\[Containments\\]\\[1\\]\\[General\\]" "$CONFIG"; then\n\
    echo "" >> "$CONFIG"\n\
    echo "[Containments][1][General]" >> "$CONFIG"\n\
    echo "url=/root/Desktop" >> "$CONFIG"\n\
  else\n\
    sed -i "/^\\[Containments\\]\\[1\\]\\[General\\]$/a url=/root/Desktop" "$CONFIG"\n\
  fi\n\
  \n\
  chown abc:abc "$CONFIG"\n\
  kquitapp6 plasmashell 2>/dev/null || pkill plasmashell\n\
  sleep 1\n\
  kstart plasmashell\n\
fi\n\
' > /usr/local/bin/fix-desktop-url.sh && \
    chmod +x /usr/local/bin/fix-desktop-url.sh && \
    printf '[Desktop Entry]\nType=Application\nName=Fix Desktop URL\nExec=/usr/local/bin/fix-desktop-url.sh\nX-KDE-autostart-phase=2\nOnlyShowIn=KDE;\n' > /root/.config/autostart/fix-desktop-url.desktop && \
    chown -R abc:abc /root/.config/autostart

# 修复 nginx fancyindex 模块版本不匹配问题 - 通过 s6 custom-init 脚本在运行时修复
RUN rm -f /etc/nginx/modules.d/fancy.conf || true && \
    mkdir -p /custom-cont-init.d && \
    echo '#!/bin/bash' > /custom-cont-init.d/99-fix-nginx && \
    echo 'sed -i "s/^\s*fancyindex/#fancyindex/g" /etc/nginx/conf.d/default.conf 2>/dev/null || true' >> /custom-cont-init.d/99-fix-nginx && \
    chmod +x /custom-cont-init.d/99-fix-nginx

ENV ANTHROPIC_BASE_URL="https://coding.dashscope.aliyuncs.com/apps/anthropic"
ENV ANTHROPIC_MODEL="glm-5"
ENV API_TIMEOUT_MS="3000000"
ENV CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"

COPY --from=builder-vsc-init /build/app /usr/local/bin/vsc-init
# Run vsc-init to install VS Code extensions
RUN /usr/local/bin/vsc-init

COPY --from=builder-shell-init /build/app /usr/local/bin/shell-init
# Run shell-init
RUN /usr/local/bin/shell-init

COPY --from=builder-vibe-init /build/app /usr/local/bin/vibe-init
# Run vibe-init
RUN /usr/local/bin/vibe-init

COPY --from=builder-entrypoint /build/app /docker-dev/entrypoint

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN go env -w GOPROXY=https://goproxy.cn,direct