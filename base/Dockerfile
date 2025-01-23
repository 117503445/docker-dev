FROM archlinux
LABEL maintainer="117503445"
ARG china_mirror

WORKDIR /scripts
COPY ./scripts/arch_init.sh .
RUN ./arch_init.sh && rm arch_init.sh

RUN mkdir -p /docker-dev
RUN cd /tmp && git clone https://github.com/117503445/docker-dev.git --depth=1 && cd docker-dev/entrypoint && go build . && cp ./entrypoint /docker-dev/entrypoint

RUN su - builder -c "yay -Su code-server --noconfirm"
ADD https://github.com/gdraheim/docker-systemctl-replacement/raw/master/files/docker/systemctl3.py /usr/local/bin/systemctl
RUN chmod +x /usr/local/bin/systemctl

EXPOSE 8080

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

WORKDIR /workspace

ENTRYPOINT [ "/docker-dev/entrypoint" ]