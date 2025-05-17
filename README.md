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

### `examples/exp4` - CUDA

开发 AI 项目时，需要使用 CUDA。

首先确保宿主机已安装 NVIDIA 驱动和 `nvidia-container-toolkit`

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    volumes:
        - ./:/workspace
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: 1
              capabilities: [gpu]
```

其他 Python 工程文件见 `exp4` 目录。容器内可使用具有 CUDA 加速的 Pytorch

```sh
docker compose exec dev uv run main.py
```

输出

```
torch.cuda.is_available: True
```

### `examples/exp5` - DinD CUDA

Dind(Docker in Docker) 同样可以使用 CUDA

```yaml
services:
  dev:
    build:
      dockerfile: dev.Dockerfile
    volumes:
        - ./:/workspace
        - ./scripts/entrypoint.sh:/entrypoint
        - docker:/var/lib/docker
    privileged: true
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: 1
              capabilities: [gpu]
volumes:
  docker:
```

```dockerfile
# dev.Dockerfile
FROM 117503445/dev

RUN pacman -Sy --noconfirm nvidia-container-toolkit
```

验证

```sh
# 1. 进入容器 Shell
docker compose exec dev zsh
# 2a. 验证 DinD 是否可以使用 CUDA
docker run --rm -it --gpus=all nvcr.io/nvidia/k8s/cuda-sample:nbody nbody -gpu -benchmark
# 2b. 使用国内镜像
docker run --rm -it --gpus=all registry.cn-hangzhou.aliyuncs.com/117503445-mirror/sync:linux.amd64.nvcr.io.nvidia.k8s.cuda-sample.nbody nbody -gpu -benchmark
```

### `examples/exp6` - 使用 Code Server 和 go-task

Code Server 是一个基于 VS Code 的浏览器端 IDE。`117503445/dev` 镜像内置了 Code Server，开发者在任何一台设备的浏览器上都可以访问开发环境。

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    restart: unless-stopped
    ports:
      - "4444:4444"
    volumes:
      - ./:/workspace
```

在 `compose.override.yaml` 中注入 `CODE_SERVER_PASSWORD` 环境变量。在实际项目中，可以将 `compose.override.yaml` 放入 `.gitignore` 中，避免 Code Server 密码泄露。

```yaml
# compose.override.yaml
services:
  dev:
    environment:
      - CODE_SERVER_PASSWORD=K8bDE57LaAp0vp
```

在浏览器中输入 `http://SERVER_IP:4444`，即可访问 Code Server。

![code server](docs/assets/1.png)

`go-task` 是一个 Go 编写的命令运行器。详情见 [go-task documentation](https://taskfile.dev/usage/)。

```yaml
# Taskfile.yml
version: '3'

tasks:
  default:
    desc: "The default task" 
    cmds:
      - clear
      - task: run
      
  run:
    cmds:
      - go run .
  
  build:
    cmds:
      - go build .
```

在 Code Server 中，按下 F5，即可在终端中运行 `go-task`，也就是 `go run .`。

### `examples/exp7` - 满血 Code Server

`exp6` 基于端口访问容器 Code Server，当宿主机上有十几个项目容器时，存在以下问题

- `http://SERVER_IP:4444` 属于 `insecure context`，无法启用剪贴板和 PWA。只有 PWA 模式下可以使用 `ctrl w` 等快捷键，因此无法启用 PWA 会导致开发体验下降。而 `http://localhost` 和 `https` 属于安全上下文，可以使用所有功能
- 端口冲突，需要给每个容器手动指定宿主机端口
- 语义不清晰，难以获悉端口号和项目的对应关系

可以使用 SSH 端口转发等方式，将 Code Server 端口映射到开发机的 localhost 上，从而启用剪贴板和 PWA。但这样仍然无法解决后面 2 个问题，而且每次需要使用某个容器时都要打开 SSH 端口转发，非常麻烦。

最理想的方式，就是容器启动后，自动生成一个 https 域名，可以访问容器内的 Code Server。比如 project0 启动后，自动就可以从 `https://vsc-project0.117503445.top` 访问开发环境。为了实现这一点，可以进行以下流程

1. 参考 [中小型应用运维](https://wiki.117503445.top/practice/中小型应用运维)，在公网云服务器上配置 Traefik 网关，支持泛域名 HTTPS 和基于 Host 的路由。
2. 参考 [traefik-provider-frp: 将 frps 代理信息提供给 Traefik，从而实现自动反向代理](https://zhuanlan.zhihu.com/p/26025560346)，在公网服务器上部署 `frps` 和 `traefik-provider-frp` 服务，允许自动为 frp 连接创建 https 子域名和路由。
3. 参考 [frpc-controller: 基于 Docker Labels 自动生成 frpc 配置文件](https://zhuanlan.zhihu.com/p/26026511814)，在运行容器的开发服务器上，运行 `frpc-controller`，自动为具有特定标签的 Docker 容器创建 frp 连接。

然后可以很简单地编写 Docker Compose

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    restart: unless-stopped
    volumes:
      - ./:/workspace
```

```yaml
services:
  dev:
    environment:
      - CODE_SERVER_PASSWORD=K8bDE57LaAp0vp
    networks:
      - frp
    labels:
      - frpc.vsc-exp7=4444

networks:
  frp:
    external: true
```

启动容器后，即可通过 https 域名访问容器，并可以启用 PWA

![pwa](docs/assets/2.png)

启用 PWA 后，Code Server 和 本地 VSCode 的体验几乎一致

## `examples/exp8` - 通义灵码

通义灵码是一款很好用的 AI 辅助编码工具，需要持久化 `/root/.lingma` 和 `/root/.cache`。

```yaml
# compose.yaml
services:
  dev:
    image: 117503445/dev
    restart: unless-stopped
    volumes:
      - ./:/workspace
      - lingma:/root/.lingma
      - cache:/root/.cache
volumes:
  lingma:
  cache:
```
