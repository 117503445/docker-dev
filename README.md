# docker-dev

use docker to build development environment

## pull

```sh
docker pull 117503445/dev-base
docker pull 117503445/dev-coding
docker pull 117503445/dev-golang
docker pull 117503445/dev-python
docker pull 117503445/dev-front
docker pull 117503445/dev-rust
docker pull 117503445/dev-cpp
docker pull 117503445/dev-java
docker pull 117503445/dev-typst
docker pull 117503445/dev-dind
docker pull 117503445/dev-csharp
docker pull 117503445/dev-latex

# China mirror
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base 117503445/dev-base
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-coding && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-coding 117503445/dev-coding
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang 117503445/dev-golang
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python 117503445/dev-python
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front 117503445/dev-front
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-rust && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-rust 117503445/dev-rust
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-cpp && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-cpp 117503445/dev-cpp
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-java && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-java 117503445/dev-java
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-typst && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-typst 117503445/dev-typst
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex 117503445/dev-latex
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-csharp && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-csharp 117503445/dev-csharp
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-dind && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-dind 117503445/dev-dind
```

## usage

```sh
docker run -it -v $PWD:/root/project 117503445/dev-front
docker run -it -v $PWD:/root/project 117503445/dev-golang
docker run -it -v $PWD:/root/project 117503445/dev-latex
docker run -it -v $PWD:/root/project 117503445/dev-python
docker run -it -v $PWD:/root/project 117503445/dev-rust
docker run -it -v $PWD:/root/project 117503445/dev-typst
docker run -it -v $PWD:/root/project 117503445/dev-cpp
docker run -it -v $PWD:/root/project 117503445/dev-java
docker run -it -v $PWD:/root/project 117503445/dev-csharp
docker run -it -v $PWD:/root/project 117503445/dev-dind
```

## local build

```sh
docker build -f ./base/Dockerfile -t 117503445/dev-base .
docker run --rm -it 117503445/dev-base

docker build -f ./golang/Dockerfile -t 117503445/dev-golang .

docker build -f ./python/Dockerfile -t 117503445/dev-python .
docker build -f ./rust/Dockerfile -t 117503445/dev-rust .
docker build -f ./typst/Dockerfile -t 117503445/dev-typst .
docker build -f ./cpp/Dockerfile -t 117503445/dev-cpp .
docker build -f ./java/Dockerfile -t 117503445/dev-java .
```

## dev container

`.devcontainer/devcontainer.json`

```jsonc
{
  "image": "117503445/dev-base",
  "customizations": {
    "vscode": {
      // Add the IDs of extensions you want installed when the container is created.
      "extensions": [
      ]
    }
  }
}
```

## misc

```sh
# install by yay
RUN su - builder -c "yay -Su scala --noconfirm"
```

## code-server

访问容器的 4444 端口即可使用 code-server

可以设置 CODE_SERVER_PASSWORD 环境变量来设置密码

/entrypoint 可以自定义启动命令