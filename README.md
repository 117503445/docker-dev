# docker-dev

use docker to build development environment

## pull

```sh
docker pull 117503445/dev-base
docker pull 117503445/dev-front
docker pull 117503445/dev-golang
docker pull 117503445/dev-latex
docker pull 117503445/dev-python

# China mirror
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-base 117503445/dev-base
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-front 117503445/dev-front
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-golang 117503445/dev-golang
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-latex 117503445/dev-latex
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev-python 117503445/dev-python
```

## usage

```sh
docker run -it -v $PWD:/root/project 117503445/dev-front
docker run -it -v $PWD:/root/project 117503445/dev-golang
docker run -it -v $PWD:/root/project 117503445/dev-latex
docker run -d --name dev-kde -p 6080:6080 -e VNC_PASSWD=password 117503445/dev-kde
docker run -it -v $PWD:/root/project 117503445/dev-python
```

## local dev

```sh
docker build -f ./base/Dockerfile -t 117503445/dev-base .
docker run --rm -it 117503445/dev-base

docker build -f ./golang/Dockerfile -t 117503445/dev-golang .
docker build -f ./kde/Dockerfile -t 117503445/dev-kde .

docker build -f ./python/Dockerfile -t 117503445/dev-python .
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
