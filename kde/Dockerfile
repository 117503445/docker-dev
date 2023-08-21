FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
COPY ./public /root/public
WORKDIR /root
RUN /root/public/kde_init.sh
EXPOSE 5900 6080
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
CMD [ "/root/public/kde_start.sh" ]