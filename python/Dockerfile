FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu poetry --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN curl -o miniconda.sh https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh && chmod +x miniconda.sh && ./miniconda.sh -b && rm ./miniconda.sh
ENV PATH="/root/miniconda3/bin:${PATH}"
RUN conda init zsh
RUN cp /root/public/.condarc /root/.condarc
ENTRYPOINT [ "/bin/zsh" ]