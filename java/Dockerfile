FROM 117503445/dev-base
LABEL maintainer="117503445"

RUN pacman -Sy jdk-openjdk maven gradle --noconfirm

RUN su - builder -c "yay -Su scala --noconfirm"

ENV VSC_EXTS="vscjava.vscode-java-pack,vscjava.vscode-gradle,fwcd.kotlin,vscjava.vscode-spring-boot-dashboard,vmware.vscode-boot-dev-pack,${VSC_EXTS}"