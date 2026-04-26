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

- 根目录 `Taskfile.yml` 只做全局配置、`includes` 和稳定入口转发，不要堆积具体业务命令。
- 领域任务放到 `scripts/tasks/<领域>/Taskfile.yml`，例如 `build`、`run`、`gen`、`deploy`、`format`、`fe`、`test`、`e2e`。
- 所有对用户可见的任务必须有中文 `desc`。
- 可复用但不希望用户直接调用的任务设置 `internal: true`。
- 跨 include 依赖使用完整命名空间，例如 `:base:clear`、`:build:bin`。
- 任务命名使用小写短横线或既有项目风格；同一仓库内保持一致。
- 保留顶层 `test` 命令作为全量测试入口，必须覆盖单元测试、集成测试和 E2E 测试。
- 保留顶层 `e2e` 命令作为 E2E 测试入口，并支持 `go-task e2e -- --case <name>` 运行单个用例。
- Taskfile 不能包含复杂逻辑，越简单越好；复杂逻辑写在 `scripts/go-scripts/` 中，再由 Taskfile 调用。
- 所有任务都要确保基于本地最新代码执行；运行、测试、E2E 等任务必须通过 `deps` 依赖必要的生成或构建任务，例如 E2E 依赖后端和前端构建。
- 编译、代码生成等无副作用任务应写好 `sources` 和 `generates`，确保代码不变时不重新执行。
- 有副作用的任务不要配置 `sources` 和 `generates`，例如 `run`、`deploy`、`test`、`e2e`。
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
  test:
    taskfile: ./scripts/tasks/test
  e2e:
    taskfile: ./scripts/tasks/e2e

tasks:
  test:
    desc: "运行所有测试"
    cmds:
      - task: test:all

  e2e:
    desc: "运行 E2E 测试"
    cmds:
      - task: e2e:all
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

### test 和 e2e 入口

项目应保留稳定的 `test` 和 `e2e` 顶层任务。`test` 是全量测试入口，必须包含 E2E；`e2e` 是 E2E 专用入口，需要支持单个 case 参数。

```yaml
tasks:
  test:
    desc: "运行所有测试"
    cmds:
      - task: test:all

  e2e:
    desc: "运行 E2E 测试"
    cmds:
      - task: e2e:all
```

`scripts/tasks/test/Taskfile.yml` 示例：

```yaml
version: 3

tasks:
  all:
    desc: "运行所有测试"
    deps:
      - ":gen:rpc"
    cmds:
      - go test ./...
      - task: :e2e:all
```

`scripts/tasks/e2e/Taskfile.yml` 示例：

```yaml
version: 3

tasks:
  all:
    desc: "运行 E2E 测试"
    deps:
      - ":build:bin"
      - ":fe:build"
    dir: ./scripts/e2e
    cmds:
      - uv run main.py {{.CLI_ARGS}}
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
    sources:
      - ./package.json
      - ./pnpm-lock.yaml
      - ./src/**
      - ./index.html
      - ./vite.config.ts
    generates:
      - ./dist/**
    cmds:
      - pnpm build

  case:
    desc: "运行指定的 E2E 测试用例"
    deps:
      - ":build:bin"
      - ":fe:build"
    dir: ./scripts/e2e
    cmds:
      - uv run main.py {{.CLI_ARGS}}
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
4. 确认顶层 `test` 和 `e2e` 任务存在；如果缺失，新增只做转发的稳定入口。
5. 为运行和测试任务补齐构建依赖，确保任务基于本地最新代码执行。
6. 为编译、代码生成等无副作用任务补齐 `sources` 和 `generates`，让 go-task 缓存判断可靠。
7. 不给有副作用的任务配置 `sources` 和 `generates`。
8. 跨领域依赖写成 `:领域:任务`，避免相对 include 解析歧义。
9. 修改后运行 `task --list` 检查任务是否可发现。
10. 按项目约定运行测试；代码变更后运行 `go-task test`，E2E 变更先运行 `go-task e2e -- --case <name>` 再运行 `go-task e2e`。

## 测试流程

如果项目可以运行 `go-task e2e`，并且本次需要修改代码：

1. 先在测试用例中覆盖新增需求。
2. 运行 `go-task e2e -- --case <name>`，预期先因为新增用例失败。
3. 修改实现代码。
4. 再运行 `go-task e2e -- --case <name>`，直到通过。
5. 运行 `go-task test`，确认所有测试通过。
6. 密钥文件必须放在 `.env` 中，并在代码中主动加载 `.env`。
7. 不应进入版本控制的生成文件、缓存、日志和构建产物必须用 `.gitignore` 排除。

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

- 运行 E2E 测试和所有测试：

```bash
go-task e2e -- --case <name>
go-task e2e
go-task test
```

## 验收标准

- `task --list` 中的任务描述为中文。
- 顶层保留可执行的 `test` 全量测试任务。
- 顶层保留可执行的 `e2e` E2E 测试任务。
- 根 `Taskfile.yml` 保持轻量，只负责 `dotenv`、`includes` 和稳定入口转发。
- 领域任务文件路径稳定，命名空间清晰。
- 依赖链可执行，不依赖当前 shell 的隐式工作目录。
- 运行、测试、E2E 任务通过 `deps` 依赖必要的构建任务，避免使用过期构建产物。
- 编译和代码生成任务配置了准确的 `sources` 和 `generates`；有副作用任务没有配置这些缓存字段。
- 项目要求的验证命令通过，例如 `go-task e2e` 和 `go-task test`。
