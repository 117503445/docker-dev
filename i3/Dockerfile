FROM 117503445/dev-base
LABEL maintainer="117503445"
# RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

COPY ./public /root/public

WORKDIR /root
RUN /root/public/i3_init.sh

EXPOSE 5900 6080
CMD [ "/root/public/i3_start.sh" ]