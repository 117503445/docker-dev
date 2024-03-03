FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu typst noto-fonts noto-fonts-cjk adobe-source-han-sans-cn-fonts noto-fonts-emoji --noconfirm
COPY typst/fonts /root/.local/share/fonts
ENV TYPST_FONT_PATHS=/root/.local/share/fonts
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN /scripts/install_vsc_ext.py nvarner.typst-lsp tomoki1207.pdf
