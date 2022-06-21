FROM 117503445/dev-base
LABEL maintainer="117503445"
ENV PATH="/root/go/bin:${PATH}"
RUN pacman -S go --noconfirm
RUN wget -O gf -q https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH) && chmod +x gf && ./gf install -y && rm ./gf
ENTRYPOINT [ "zsh" ]