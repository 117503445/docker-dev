FROM 117503445/dev-base
LABEL maintainer="117503445"

WORKDIR /scripts
COPY ./scripts/set_go.sh .
RUN ./set_go.sh && rm set_go.sh
WORKDIR /workspace

# ENV VSC_EXTS="golang.go,${VSC_EXTS}"