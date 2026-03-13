ARG BASE_IMAGE=lscr.io/linuxserver/webtop:arch-kde

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

# 修复 nginx fancyindex 模块版本不匹配问题 - 通过 s6 custom-init 脚本在运行时修复
RUN rm -f /etc/nginx/modules.d/fancy.conf || true && \
    mkdir -p /custom-cont-init.d && \
    echo '#!/bin/bash' > /custom-cont-init.d/99-fix-nginx && \
    echo 'sed -i "s/^\s*fancyindex/#fancyindex/g" /etc/nginx/conf.d/default.conf 2>/dev/null || true' >> /custom-cont-init.d/99-fix-nginx && \
    chmod +x /custom-cont-init.d/99-fix-nginx

# 安装所需软件包
RUN pacman -Syu which zsh fish btop git openssh docker docker-compose docker-buildx nano vim micro base-devel parted tmux python wget yazi go-task go zsh-autosuggestions zsh-syntax-highlighting podman skopeo umoci ncdu cloc fuse-overlayfs nvidia-container-toolkit github-cli jq rsync kubectl helm unzip --noconfirm

# 添加用户 builder，并安装 yay AUR 助手
RUN useradd -m -G wheel builder && \
    chown -R builder:builder /tmp && \
    echo 'builder ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    su - builder -c "cd /tmp && git clone https://aur.archlinux.org/yay-bin.git && cd yay-bin && makepkg -si --noconfirm" && \
    rm -rf /tmp/yay

ENV PATH="$HOME/go/bin:${PATH}"
# # set gopath
# RUN go env -w GOPROXY=direct
RUN go env -w GOPATH=$HOME/go
RUN go install github.com/117503445/histmon@master
RUN su - builder -c "yay -Su code-server claude-code cursor-cli openai-codex-bin --noconfirm"
RUN GOPROXY=direct go install github.com/117503445/vsc-init@master
RUN $HOME/go/bin/vsc-init
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

# 更改默认 shell 为 zsh
RUN chsh -s /usr/bin/zsh
# RUN su - builder -c "yay -Su --noconfirm oh-my-zsh-git"
RUN if [ -z "$CHINA_MIRROR" ]; then \
      sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"; \
    else \
      sh -c "$(curl -fsSL https://install.ohmyz.sh)"; \
    fi
COPY ./assets/.zshrc /tmp/.zshrc
RUN mv /tmp/.zshrc $HOME/.zshrc
RUN touch $HOME/.zshrc-custom
RUN TOKEN="T" ENDPOINT="E" histmon install

RUN go install github.com/cweill/gotests/gotests@latest && go install github.com/fatih/gomodifytags@latest && go install github.com/josharian/impl@latest && go install github.com/haya14busa/goplay/cmd/goplay@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install honnef.co/go/tools/cmd/staticcheck@latest && go install golang.org/x/tools/gopls@latest && go install github.com/ramya-rao-a/go-outline@latest && go install github.com/goreleaser/goreleaser/v2@latest && go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 && go install -v github.com/incu6us/goimports-reviser/v3@latest && go install github.com/bufbuild/buf/cmd/buf@latest && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest && rm -rf $HOME/.cache $HOME/go/pkg

# front
RUN pacman -Syu npm yarn pnpm --noconfirm
RUN npm config set registry https://registry.npmmirror.com
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn global add typescript
RUN pnpm config set registry https://registry.npmmirror.com

# python
RUN pacman -Syu uv --noconfirm
COPY ./assets/python/uv.toml /tmp/uv.toml
RUN mkdir -p $HOME/.config/uv && mv /tmp/uv.toml $HOME/.config/uv/uv.toml

# latex
# RUN pacman -Syu texlive make gnu-free-fonts texlive-langchinese biber --noconfirm
RUN git clone https://github.com/Haixing-Hu/latex-chinese-fonts.git /usr/share/fonts/latex-chinese-fonts --depth=1

# cpp
RUN pacman -Syu gcc clang cmake --noconfirm

# csharp
RUN pacman -Syu --noconfirm dotnet-runtime dotnet-sdk aspnet-runtime

# java
RUN pacman -Sy jdk-openjdk maven gradle kotlin --noconfirm
# RUN su - builder -c "yay -Su scala --noconfirm"

# # rust
RUN pacman -Syu rust-analyzer --noconfirm
# RUN sh -c "if [ -n \"$CHINA_MIRROR\" ]; then \
#                 export RUSTUP_DIST_SERVER=https://rsproxy.cn; \
#                 export RUSTUP_UPDATE_ROOT=https://rsproxy.cn/rustup; \
#                 rustup default stable; \
#             else \
#                 rustup default stable; \
#             fi"
# RUN rustup default stable
RUN cargo install cargo-watch
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"
ENV PATH="$HOME/.cargo/bin:${PATH}"
COPY ./assets/rust/config.toml /tmp/cargo-config.toml
RUN mkdir -p $HOME/.cargo && mv /tmp/cargo-config.toml $HOME/.cargo/config.toml

# typst
RUN pacman -Syu typst noto-fonts noto-fonts-cjk adobe-source-han-sans-cn-fonts noto-fonts-emoji ttf-ubuntu-font-family --noconfirm
# COPY typst/fonts /usr/share/fonts/docker-dev
ENV TYPST_FONT_PATHS=/usr/share/fonts
RUN wget https://github.com/Enter-tainer/typstyle/releases/latest/download/typstyle-x86_64-unknown-linux-gnu -o typstyle && chmod +x typstyle && mv typstyle /usr/bin/typstyle

# aliyun
RUN npm install @serverless-devs/s -g
RUN wget https://gosspublic.alicdn.com/ossutil/1.7.19/ossutil-v1.7.19-linux-amd64.zip && unzip ossutil-v1.7.19-linux-amd64.zip && mv ossutil-v1.7.19-linux-amd64/ossutil /usr/local/bin/ && rm ossutil-v1.7.19-linux-amd64.zip ossutil-v1.7.19-linux-amd64/ossutil64 && rm -r ossutil-v1.7.19-linux-amd64

RUN mkdir -p /docker-dev/logs

WORKDIR /workspace

RUN curl https://wiki.117503445.top/linux/script/ssh.sh | bash
COPY ./assets/sshd/ssh_host_ed25519_key /etc/ssh/ssh_host_ed25519_key
COPY ./assets/sshd/ssh_host_ed25519_key.pub /etc/ssh/ssh_host_ed25519_key.pub
RUN chmod 600 /etc/ssh/ssh_host_ed25519_key

ENV TZ=Asia/Shanghai
ENV DOCKER_BUILDKIT=1
ENV PATH="$HOME/.local/bin:${PATH}"

ENV ANTHROPIC_BASE_URL="https://coding.dashscope.aliyuncs.com/apps/anthropic"
ENV ANTHROPIC_MODEL="glm-5"
ENV API_TIMEOUT_MS="3000000"
ENV CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"
COPY ./assets/claude/settings.json /tmp/claude-settings.json
RUN mkdir -p $HOME/.claude && mv /tmp/claude-settings.json $HOME/.claude/settings.json

COPY ./scripts/entrypoint/entrypoint /docker-dev/entrypoint

# 配置 custom service 启动 /docker-dev/entrypoint (linuxserver 镜像推荐方式)
RUN mkdir -p /custom-services.d && \
    printf '#!/usr/bin/with-contenv bash\nexec /docker-dev/entrypoint\n' > /custom-services.d/docker-dev-entrypoint && \
    chmod +x /custom-services.d/docker-dev-entrypoint

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 创建 KDE 桌面模板文件，使右键菜单可以创建新文件
RUN mkdir -p $HOME/Templates && \
    touch "$HOME/Templates/Text File.txt" && \
    touch "$HOME/Templates/Empty File"

# 配置 abc 用户以 root 权限运行桌面，并设置桌面路径为 /root/Desktop
# LinuxServer 镜像使用 abc 用户运行桌面环境，需要修改其 home 目录
RUN usermod -d /root abc && \
    mkdir -p /root/Desktop && \
    chown -R abc:abc /root && \
    chmod 755 /root /root/Desktop

# 配置 KDE Plasma 桌面 Folder View 显示 /root/Desktop 内容
# 创建 KDE autostart 脚本，在桌面启动后修复配置
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
