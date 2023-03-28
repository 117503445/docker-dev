# kde

dev

```sh
docker rm -f dev-kde
docker build -f ./kde/Dockerfile -t 117503445/dev-kde .
docker run -d --name dev-kde -p 6080:6080 -e VNC_PASSWD=password -v ${PWD}/public:/root/public 117503445/dev-kde

docker rm -f dev-kde
docker run -d --name dev-kde -p 6080:6080 -e VNC_PASSWD=password -v ${PWD}/public:/root/public 117503445/dev-kde

docker rm -f dev-kde
docker run -d --name dev-kde -p 6080:6080 -e VNC_PASSWD=password -e KDE_USERNAME=htqi -v ${PWD}/public:/root/public 117503445/dev-kde

docker logs -f dev-kde

docker exec -it 117503445/dev-kde /bin/bash
```
