FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu texlive make gnu-free-fonts texlive-langchinese biber --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

RUN git clone https://github.com/Haixing-Hu/latex-chinese-fonts.git /usr/share/fonts/latex-chinese-fonts --depth=1