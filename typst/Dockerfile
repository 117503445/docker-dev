FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu typst noto-fonts noto-fonts-cjk adobe-source-han-sans-cn-fonts noto-fonts-emoji ttf-ubuntu-font-family --noconfirm
COPY typst/fonts /usr/share/fonts/docker-dev
ENV TYPST_FONT_PATHS=/usr/share/fonts
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN wget https://github.com/Enter-tainer/typstyle/releases/latest/download/typstyle-x86_64-unknown-linux-gnu -o typstyle && chmod +x typstyle && mv typstyle /usr/bin/typstyle
ENV VSC_EXTS="nvarner.typst-lsp,tomoki1207.pdf,${VSC_EXTS}"