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
- 领域任务放到 `scripts/tasks/<领域>/Taskfile.yml`，例如 `build`、`run`、`gen`、`deploy`、`format`、`fe`、`test`。
- 所有对用户可见的任务必须有中文 `desc`。
- 可复用但不希望用户直接调用的任务设置 `internal: true`。
- 跨 include 依赖使用完整命名空间，例如 `:base:clear`、`:build:bin`。
- 任务命名使用小写短横线或既有项目风格；同一仓库内保持一致。
- 保留顶层 `run` 命令作为本地运行入口，默认运行完整后端服务，并依赖必要构建任务。
- 保留顶层 `test` 命令作为全量测试入口，必须覆盖单元测试、集成测试和 E2E 测试。
- `scripts/tasks/test/Taskfile.yml` 必须提供 `ut`、`it`、`e2e` 和 `all` 任务，对应 `go-task test:ut`、`go-task test:it`、`go-task test:e2e` 和 `go-task test`。
- `test:ut` 的日志必须输出到 `./data/ut/`；`test:it` 的日志必须输出到 `./data/it/`。
- 集成测试指先编译并启动一个 Go 服务，再使用 Go client 调用该服务验证行为；启动、等待、调用和清理逻辑写在 `scripts/go-scripts it` 中。
- 每个 IT/E2E case 都必须自己启动一套新的服务，运行测试代码，把服务日志输出到自己的 case 目录，最后关闭这套服务，避免 case 间共享服务状态。
- IT/E2E 的 `server.log` 不要输出 ASCII/ANSI 颜色控制字符；启动服务时需要传入 `nocolor` 参数，并让 zerolog console writer 禁用颜色。
- Taskfile 不能包含复杂逻辑，越简单越好；复杂逻辑写在 `scripts/go-scripts/` 中，再由 Taskfile 调用。
- `scripts/go-scripts` 根包只负责 CLI 解析、命令注册和分发；每个命令的实际实现必须放到子模块（Go 子包），例如 `scripts/go-scripts/internal/build`、`scripts/go-scripts/internal/it`、`scripts/go-scripts/internal/e2e`，不要把 `build.go`、`release.go` 这类实现文件直接放在根包。
- E2E 如需浏览器自动化或 Playwright，必须使用 Go 和 `github.com/playwright-community/playwright-go` 实现，不使用 Python Playwright。
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

tasks:
  run:
    desc: "运行本地服务"
    cmds:
      - task: run:rpc

  test:
    desc: "运行所有测试"
    cmds:
      - task: test:all
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
      - go run ./scripts/go-scripts build
```

`scripts/go-scripts` 命令实现结构：

```text
scripts/go-scripts/
├── main.go                    # 程序入口，只初始化日志和解析 CLI
├── cli.go                     # 命令定义和分发，不写具体业务逻辑
└── internal/
    ├── build/                 # build 命令实现
    ├── release/               # release 命令实现
    ├── deploy/                # deploy 命令实现
    ├── docker/                # build-docker 命令实现
    ├── format/                # format 命令实现
    ├── invoke/                # invoke 命令实现
    ├── it/                    # it 命令实现
    └── e2e/                   # e2e 命令实现
```

## 常用模式

### run 入口

项目应保留稳定的 `run` 顶层任务，作为本地运行入口。`run` 本身只做转发，具体运行逻辑放到 `scripts/tasks/run/Taskfile.yml`。

```yaml
tasks:
  run:
    desc: "运行本地服务"
    cmds:
      - task: run:rpc
```

`scripts/tasks/run/Taskfile.yml` 示例：

```yaml
version: 3

tasks:
  rpc:
    desc: "运行 RPC 服务"
    deps:
      - ":build:bin"
    cmds:
      - ./data/rpc/rpc {{.CLI_ARGS}}
```

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

### test 入口

项目应保留稳定的 `test` 顶层任务，作为全量测试入口。单元测试、集成测试和 E2E 都放在 `test` 命名空间下，不提供独立 E2E 顶层入口。

```yaml
tasks:
  test:
    desc: "运行所有测试"
    cmds:
      - task: test:all
```

`scripts/tasks/test/Taskfile.yml` 示例：

```yaml
version: 3

tasks:
  all:
    desc: "运行所有测试"
    cmds:
      - task: ut
      - task: it
      - task: e2e

  ut:
    desc: "运行单元测试"
    deps:
      - ":gen:rpc"
    cmds:
      - mkdir -p ./data/ut
      - go test ./... -short 2>&1 | tee ./data/ut/test.log

  it:
    desc: "运行集成测试"
    deps:
      - ":build:bin"
    cmds:
      - mkdir -p ./data/it
      - go run ./scripts/go-scripts it 2>&1 | tee ./data/it/test.log

  e2e:
    desc: "运行 E2E 测试"
    deps:
      - ":build:bin"
      - ":fe:build"
    cmds:
      - go run ./scripts/go-scripts e2e {{.CLI_ARGS}}
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
      - go run ./scripts/go-scripts build-docker {{.CLI_ARGS}}
```

调用示例：

```bash
task build:docker -- --push
```

### 集成测试脚本

`test:it` 只负责调用 `scripts/go-scripts it` 并收集日志。`scripts/go-scripts it` 内部负责按 case 启动已编译的 Go 服务、等待健康检查通过、使用 Go client 调用服务接口验证行为，并在每个 case 结束时清理服务进程。每个 case 的服务日志写入 `./data/it/<case_name>/server.log`。启动服务时传入 `nocolor` 参数，使服务端 zerolog console writer 禁用颜色，避免 `server.log` 写入 ASCII/ANSI 颜色控制字符。

`test:e2e` 只负责调用 `scripts/go-scripts e2e` 并收集结果。`scripts/go-scripts/internal/e2e` 内部每个 case 负责启动一套新的服务，使用 `github.com/playwright-community/playwright-go` 运行浏览器测试，把服务日志写入 `./data/e2e/<case_name>/logs/server.log`，并在浏览器测试结束后清理服务进程。启动服务时同样传入 `nocolor` 参数，确保 `server.log` 不包含颜色控制字符。首次运行或 CI 初始化时，按 `go.mod` 中的 playwright-go 版本安装浏览器驱动，或在代码中显式调用 `playwright.Install()`。

### 指定工作目录

前端等子项目优先使用 `dir`，避免在命令里反复 `cd`；E2E 入口统一通过 `scripts/go-scripts e2e` 分发：

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

  e2e:
    desc: "运行指定的 E2E 测试用例"
    deps:
      - ":build:bin"
      - ":fe:build"
    cmds:
      - go run ./scripts/go-scripts e2e {{.CLI_ARGS}}
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
4. 确认顶层只保留 `run` 和 `test` 稳定入口。
5. 为运行和测试任务补齐构建依赖，确保任务基于本地最新代码执行。
6. 为编译、代码生成等无副作用任务补齐 `sources` 和 `generates`，让 go-task 缓存判断可靠。
7. 不给有副作用的任务配置 `sources` 和 `generates`。
8. 跨领域依赖写成 `:领域:任务`，避免相对 include 解析歧义。
9. 修改后运行 `task --list` 检查任务是否可发现。
10. 按项目约定运行测试；代码变更后运行 `go-task test`，E2E 变更可先运行 `go-task test:e2e -- --case <name>` 调试单个用例。

## 测试流程

如果项目可以运行 `go-task test`

1. 根据变更风险判断是否需要新增测试用例；影响用户流程或存在回归风险时补充 E2E，否则优先使用 UT、IT 或已有 E2E 覆盖。
2. 新增 E2E 时先运行 `go-task test:e2e -- --case <name>`，预期先因为新增用例失败。
3. 修改实现代码。
4. 如果新增或修改了 E2E，再运行 `go-task test:e2e -- --case <name>`，直到通过。
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

- 运行本地服务和测试：

```bash
go-task run
go-task test:ut
go-task test:it
go-task test:e2e -- --case <name>
go-task test
```

## 验收标准

- `task --list` 中的任务描述为中文。
- 顶层保留可执行的 `run` 本地运行任务。
- 顶层保留可执行的 `test` 全量测试任务。
- 根 `Taskfile.yml` 保持轻量，只负责 `dotenv`、`includes` 和稳定入口转发。
- 领域任务文件路径稳定，命名空间清晰。
- 依赖链可执行，不依赖当前 shell 的隐式工作目录。
- 运行、测试、E2E 任务通过 `deps` 依赖必要的构建任务，避免使用过期构建产物。
- 编译和代码生成任务配置了准确的 `sources` 和 `generates`；有副作用任务没有配置这些缓存字段。
- 项目要求的验证命令通过，例如 `go-task run` 和 `go-task test`。
