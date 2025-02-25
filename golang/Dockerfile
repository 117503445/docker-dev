FROM 117503445/dev-base
LABEL maintainer="117503445"

WORKDIR /scripts
RUN go env -w GOPROXY=direct
RUN go install github.com/cweill/gotests/gotests@latest && go install github.com/fatih/gomodifytags@latest && go install github.com/josharian/impl@latest && go install github.com/haya14busa/goplay/cmd/goplay@latest && go install github.com/go-delve/delve/cmd/dlv@latest && go install honnef.co/go/tools/cmd/staticcheck@latest && go install golang.org/x/tools/gopls@latest && go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/goreleaser/goreleaser/v2@latest
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /workspace