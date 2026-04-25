---
name: go-task
description: |
  维护 go-task / Taskfile 任务体系。适用于新增、重构或排查 Taskfile.yml，
  设计构建、运行、生成、部署、前端、E2E 等任务，以及把脚本编排接入 task 命令。
---

# go-task 任务编排技能

用于在项目中维护基于 go-task 的任务体系。目标是让常用开发命令可发现、可组合、可缓存，并和项目目录结构保持一致。

## 适用场景

- 新增或调整 `Taskfile.yml`。
- 把构建、运行、代码生成、部署、格式化、前端或 E2E 流程接入 `task`。
- 排查某个命令、脚本或文件被哪个 task 引用。
- 统一 task 描述、命名、依赖和目录约定。

## 基本原则

- 根目录 `Taskfile.yml` 只做全局配置和 `includes`，不要堆积具体业务命令。
- 领域任务放到 `scripts/tasks/<领域>/Taskfile.yml`，例如 `build`、`run`、`gen`、`deploy`、`format`、`fe`、`e2e`。
- 所有对用户可见的任务必须有中文 `desc`。
- 可复用但不希望用户直接调用的任务设置 `internal: true`。
- 跨 include 依赖使用完整命名空间，例如 `:base:clear`、`:build:bin`。
- 任务命名使用小写短横线或既有项目风格；同一仓库内保持一致。
- 保留顶层 `smoke` 命令作为冒烟测试入口；如果项目已有 `go-task smoke`，不要删除、改名或绕开它。
- 优先让 `task --list` 能清楚展示常用入口。

## 推荐结构

根目录 `Taskfile.yml`：

```yaml
version: 3

dotenv: [".env"]

includes:
  base:
    taskfile: ./scripts/tasks/base
  build:
    taskfile: ./scripts/tasks/build
  run:
    taskfile: ./scripts/tasks/run
  gen:
    taskfile: ./scripts/tasks/gen
  deploy:
    taskfile: ./scripts/tasks/deploy
  format:
    taskfile: ./scripts/tasks/format
  fe:
    taskfile: ./scripts/tasks/fe
  e2e:
    taskfile: ./scripts/tasks/e2e
```

领域任务文件 `scripts/tasks/build/Taskfile.yml`：

```yaml
version: 3

tasks:
  bin:
    desc: "构建二进制文件"
    deps: [":base:clear"]
    sources:
      - ./cmd/**
      - ./internal/**
      - ./pkg/**
      - ./go.mod
      - ./go.sum
      - ./Taskfile.yml
      - ./scripts/go-scripts/**
      - ./scripts/tasks/**
    generates:
      - ./data/cli/cli
      - ./data/rpc/rpc
    cmds:
      - cd ./scripts/go-scripts && go run . build
```

## 常用模式

### 清屏和默认任务

```yaml
version: 3

tasks:
  clear:
    internal: true
    silent: true
    run: once
    cmds:
      - clear || true

  default:
    desc: "默认任务"
    deps: ["clear"]
    cmds:
      - task: run:cli-run
```

### 冒烟测试入口

项目应保留一个稳定的 `smoke` 任务，作为代理和开发者修改代码后的快速验收入口。`smoke` 可以依赖测试、构建或关键 E2E 子集，但不要把它做成只打印信息的空任务。

```yaml
tasks:
  smoke:
    desc: "运行冒烟测试"
    cmds:
      - task: test
```

如果冒烟测试放在独立 include 中，根任务仍应保留 `smoke` 转发入口：

```yaml
tasks:
  smoke:
    desc: "运行冒烟测试"
    cmds:
      - task: test:smoke
```

### 运行构建产物

```yaml
tasks:
  cli-run:
    desc: "运行 CLI 程序"
    deps: [":base:clear", ":build:bin"]
    cmds:
      - ./data/cli/cli test

  rpc-run:
    desc: "运行 RPC 服务"
    deps: [":base:clear", ":build:bin"]
    cmds:
      - ./data/rpc/rpc
```

### 传递命令行参数

需要把用户传给 task 的剩余参数交给脚本时，使用 `{{.CLI_ARGS}}`：

```yaml
tasks:
  docker:
    desc: "构建容器镜像"
    deps: [":base:clear", ":build:bin"]
    cmds:
      - cd ./scripts/go-scripts && go run . build-docker {{.CLI_ARGS}}
```

调用示例：

```bash
task build:docker -- --push
```

### 指定工作目录

前端、E2E 等子项目优先使用 `dir`，避免在命令里反复 `cd`：

```yaml
tasks:
  build:
    desc: "构建前端"
    dir: ./fe
    cmds:
      - pnpm build

  case:
    desc: "运行指定的 E2E 测试用例"
    dir: ./scripts/e2e
    cmds:
      - uv run main.py --case {{.CLI_ARGS}}
```

### 临时补充环境变量

需要使用本地工具链时，在任务内设置 `env`：

```yaml
tasks:
  gen:
    desc: "为前端生成 protobuf 代码"
    deps: [":base:clear"]
    env:
      PATH: "{{.PATH}}:{{.PWD}}/fe/node_modules/.bin"
    cmds:
      - buf generate
```

## 修改流程

1. 先用 `rg -n "任务名|脚本名|文件名" Taskfile.yml scripts/tasks` 查清现有引用。
2. 判断是否需要新增 include；如果是新领域，创建 `scripts/tasks/<领域>/Taskfile.yml` 并在根 `Taskfile.yml` 注册。
3. 为每个用户可见任务补中文 `desc`。
4. 确认 `smoke` 任务仍然存在；如果项目没有冒烟入口，新增一个可执行的 `smoke` 任务。
5. 为构建类任务补齐 `sources` 和 `generates`，让 go-task 缓存判断可靠。
6. 跨领域依赖写成 `:领域:任务`，避免相对 include 解析歧义。
7. 修改后运行 `task --list` 检查任务是否可发现。
8. 按项目约定运行测试；如果存在 `go-task smoke`，代码变更时优先运行它，其次运行 `task test` 和受影响任务。

## 冒烟测试流程

如果项目通过 `go-task smoke` 运行冒烟测试，并且本次需要修改代码：

1. 先在测试用例中覆盖新增需求。
2. 运行 `go-task smoke`，预期先因为新增用例失败。
3. 修改实现代码。
4. 再运行 `go-task smoke`，直到通过。
5. 密钥文件必须放在 `.env` 中，并在代码中主动加载 `.env`。
6. 不应进入版本控制的生成文件、缓存、日志和构建产物必须用 `.gitignore` 排除。

## 排查技巧

- 查某个文件被哪个 task 使用：

```bash
rg -n "目标文件名|脚本名|命令片段" Taskfile.yml scripts/tasks
```

- 查看 task 展示结果：

```bash
task --list
```

- 预览某个任务的执行计划：

```bash
task --dry <任务名>
```

- 跳过缓存强制执行：

```bash
task --force <任务名>
```

- 运行冒烟测试：

```bash
go-task smoke
```

## 验收标准

- `task --list` 中的任务描述为中文。
- 顶层保留可执行的 `smoke` 冒烟测试任务。
- 根 `Taskfile.yml` 保持轻量，只负责 `dotenv`、`includes` 等全局配置。
- 领域任务文件路径稳定，命名空间清晰。
- 依赖链可执行，不依赖当前 shell 的隐式工作目录。
- 项目要求的验证命令通过，例如 `go-task smoke` 和 `task test`。
