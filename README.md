# docker-dev

use docker to build development environment

## usage

```sh
docker run -it -v $PWD:/root/project 117503445/dev-front
docker run -it -v $PWD:/root/project 117503445/dev-golang
docker run -it -v $PWD:/root/project 117503445/dev-latex
```

## local dev

```sh
docker build -f ./base/Dockerfile -t 117503445/dev-base .
docker run --rm -it 117503445/dev-base

docker build -f ./golang/Dockerfile -t 117503445/dev-golang .
```
