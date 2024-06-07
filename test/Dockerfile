FROM archlinux

# COPY ./test/entrypoint.sh /entrypoint.sh

RUN pacman -Syu fish --noconfirm

COPY ./test/entrypoint /docker-dev/entrypoint

WORKDIR /workspace

ENTRYPOINT [ "/docker-dev/entrypoint" ]