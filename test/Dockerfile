FROM archlinux

RUN pacman -Sy python --noconfirm

COPY ./public/download_vscode_server.sh ./public/download_vscode_server.sh
RUN ./public/download_vscode_server.sh

COPY ./public/install_vsc_ext.py ./public/install_vsc_ext.py
RUN /scripts/install_vsc_ext.py ms-python.python

ENTRYPOINT [ "bash" ]