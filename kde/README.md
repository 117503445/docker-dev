# kde

dev

```sh
docker rm -f dev-kde
docker build -f ./kde/Dockerfile -t 117503445/dev-kde .
docker run --rm -it --name dev-kde -v $PWD/kde/entrypoint.sh:/workspace/container/entrypoint.sh -p 6080:6080 117503445/dev-kde
```


<https://github.com/fennerm/arch-i3-novnc-docker>

<https://github.com/DCsunset/docker-archlinux-vnc>
