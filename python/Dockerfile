FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu poetry mypy rye --noconfirm
RUN poetry config virtualenvs.in-project true
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

RUN su - builder -c "yay -Su miniconda3 --noconfirm"

# ENV PATH="/root/miniconda3/bin:${PATH}"
# RUN conda init fish
COPY ./python/.condarc /root/.condarc
ENV VSC_EXTS="ms-python.black-formatter,formulahendry.code-runner,ms-python.isort,matangover.mypy,ms-python.vscode-pylance,ms-python.python,donjayamanne.python-environment-manager,${VSC_EXTS}"