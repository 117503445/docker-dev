# docker-dev

use docker to build development environment

```sh
docker build -f ./base/Dockerfile -t 117503445/dev-base .
docker run --rm -it 117503445/dev-base

docker build -f ./golang/Dockerfile -t 117503445/dev-golang .
```
