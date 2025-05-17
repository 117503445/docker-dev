# docker-dev

适用于开发的 Docker 镜像

## 快速开始

拉取镜像

```sh
docker pull 117503445/dev

# China mirror
docker pull registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev && docker image tag registry.cn-hangzhou.aliyuncs.com/117503445-mirror/dev 117503445/dev
```

运行容器

```sh
docker run -it -v $PWD:/workspace 117503445/dev
```

## 思想

开发环境构建往往是一个烦人的问题。

常规的做法是直接在系统上安装各种开发工具和环境，但这样存在以下缺点

- 缺乏隔离性，系统中
- 缺乏可复现性。同样的步骤
- 

## 高级

### 安装软件包

临时安装软件包

```sh
# install by pacman
pacman -Sy --noconfirm go

# install by yay
su - builder -c "yay -Su scala --noconfirm"
```

但是更推荐将软件包安装写进 Dockerfile 中

### code-server

访问容器的 4444 端口即可使用 code-server，在 WebIDE 中进行开发。

可以设置 `CODE_SERVER_PASSWORD` 环境变量来设置密码。

讲脚本映射到容器的 `/entrypoint` 可以自定义启动命令

## 实现

在 `dev/Dockerfile` 中定义了镜像的构建步骤。其中

  - Code Server 的配置、插件安装见 [vsc-init](https://github.com/117503445/vsc-init)
  - Entrypoint 见 `./entrypoint/main.go`，其中流程为
    - 修改 Code Server 源代码，以优化 PWA 下的体验
    - 将 `CODE_SERVER_PASSWORD` 环境变量写入 Code Server 的配置文件
    - 在后台启动 Code Server
    - 如果以 `--it` 启动容器，则进入 Shell；否则调用 `tail -f /dev/null` 防止容器退出

在 `.github/workflows/dev.yml` 中使用 GitHub Actions 定时构建镜像

## 例子

在 examples 目录下，提供了一些例子

### `examples/exp0` - 使用 Docker Compose

相比 `docker run`，建议使用 Docker Compose

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    restart: unless-stopped
    volumes:
      - ./:/workspace
```

常用操作

```sh
docker compose up -d # 启动容器
docker compose exec dev zsh # 进入容器 Shell
```

### `examples/exp1` - 自定义镜像

每个项目往往会有自己的需求，可以自定义镜像

```yaml
# compose.yaml
services:
  dev:
    build:
      dockerfile: dev.Dockerfile
    restart: unless-stopped
    volumes:
      - ./:/workspace
```

在 `dev.Dockerfile` 中添加进一步的配置

```dockerfile
# dev.Dockerfile
FROM 117503445/dev

RUN pacman -Sy --noconfirm go
```

### `examples/exp2` - SHA256 锁定镜像

`117503445/dev` 镜像每天都会更新，安装最新的软件。为了防止新版本软件与老项目不兼容，可以将项目开发时所使用的镜像进行锁定。

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev@sha256:342935c6d952acaef3149f4219761a9d433356172a9e7c9263e31bfb9f8aef7c
    restart: unless-stopped
    volumes:
      - ./:/workspace
```

### `examples/exp3` - Docker in Docker

有时候项目开发依赖于 Docker，为了提供开发环境，可以在 Docker 容器内通过挂载 docker.sock 来访问宿主机的 Docker 守护进程。但这会带来一些问题，比如使用 `-v` 参数进行卷映射时，路径是宿主机的而不是容器的。此外，不可避免的依赖于宿主的 Docker 环境，可能会有潜在的风险。

所以，`docker in docker` 是一个更加优雅的方案。在容器内运行 dockerd，防止影响到宿主机的 Docker 环境。

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    volumes:
        - ./:/workspace
        - ./scripts/entrypoint.sh:/entrypoint
        - docker:/var/lib/docker
    privileged: true
volumes:
  docker:
```

```sh
#!/usr/bin/env sh
# ./scripts/entrypoint.sh

dockerd
```

常用操作

```sh
docker compose exec dev docker info # 查看容器内的 Docker 守护进程信息
```

### `examples/exp4` - GPU

开发 AI 项目时，需要使用 GPU。

首先确保宿主机已安装 NVIDIA 驱动和 `nvidia-container-toolkit`



### `examples/exp999` - 使用 Code Server

Code Server 是一个基于 VS Code 的浏览器端 IDE。`117503445/dev` 镜像内置了 Code Server，开发者在任何一台设备的浏览器上都可以访问开发环境。

```yaml
# compose.yaml
