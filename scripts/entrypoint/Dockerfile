FROM golang:1.23.2

WORKDIR /build

# RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o app .

ENTRYPOINT ["/workspace/build.sh"]