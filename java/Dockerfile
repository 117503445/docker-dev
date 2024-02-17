FROM 117503445/dev-base
LABEL maintainer="117503445"

RUN pacman -Sy jdk-openjdk maven gradle --noconfirm

RUN /scripts/install_vsc_ext.py vscjava.vscode-java-pack vscjava.vscode-gradle fwcd.kotlin