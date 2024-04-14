# docker-dev

use docker to build development environment

## pull

```sh
docker pull 117503445/dev-base
docker pull 117503445/dev-front
docker pull 117503445/dev-golang
docker pull 117503445/dev-python
docker pull 117503445/dev-rust
docker pull 117503445/dev-cpp
docker pull 117503445/dev-java
docker pull 117503445/dev-typst
docker pull 117503445/dev-latex

# China mirror
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base 117503445/dev-base
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front 117503445/dev-front
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang 117503445/dev-golang
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python 117503445/dev-python
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-rust && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-rust 117503445/dev-rust
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-cpp && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-cpp 117503445/dev-cpp
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-java && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-java 117503445/dev-java
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-typst && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-typst 117503445/dev-typst
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex 117503445/dev-latex
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
# install vscode extensions
RUN /scripts/install_vsc_ext.py vscjava.vscode-java-pack vscjava.vscode-gradle fwcd.kotlin vscjava.vscode-spring-boot-dashboard vmware.vscode-boot-dev-pack

# install by yay
RUN su - builder -c "yay -Su scala --noconfirm"
```