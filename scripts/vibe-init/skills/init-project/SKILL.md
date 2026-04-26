---
name: init-project
description: |
  Initialize a new full-stack project following the go-template architecture patterns.
  Use this skill when creating new Go + React projects, setting up RPC services with Connect,
  configuring frontend with Vite/Tailwind, or establishing project structure with cmd/internal/pkg conventions.
  Trigger when user asks to "create a new project", "initialize a project", "set up a Go web service",
  or mentions "template project", "project scaffolding", or similar project initialization tasks.
---

# 初始化项目技能 (Init Project Skill)

按照 go-template 架构模式初始化一个新项目。本技能提供了一套全面的指南，用于设置基于 Connect RPC 的 Go 后端服务、基于 Vite/Tailwind 的 React 前端，以及相关的工具链配置。

**下载代码** 下载代码 https://github.com/117503445/templates 后，结合 templates 代码和以下指南，进行项目初始化。

---

## 项目约束

- 文档和注释使用中文；每个方法需要有中文注释，代码块内只在必要处添加精简注释。
- 前端使用 TypeScript 编写；没有特殊声明时，其他代码和脚本使用 Go 编写。
- 测试入口统一使用 `go-task test`；该任务必须运行所有测试，包括单元测试、集成测试和 E2E 测试。
- 不保留独立 E2E 顶层入口；单独调试 E2E 时使用 `go-task test:e2e -- --case <name>`。
- 如果需要修改代码，并且项目可以运行 `go-task test`，先在测试用例中覆盖新增需求；E2E 变更先运行 `go-task test:e2e -- --case <name>` 并确认失败，实现后继续运行该命令直到成功，最后运行 `go-task test`。
- `test:ut` 的日志必须输出到 `./data/ut/`；`test:it` 的日志必须输出到 `./data/it/`。
- 集成测试指先编译并启动一个 Go 服务，再使用 Go client 调用该服务验证行为；启动、等待、调用和清理逻辑写在 `scripts/go-scripts it` 中。
- 本地运行入口统一使用 `go-task run`。
- 修改后必须运行 `go-task test`，并保证通过。
- Taskfile 不能包含复杂逻辑，越简单越好；如果需要复杂逻辑，写在 `scripts/go-scripts/` 中，再由 Taskfile 调用。
- 密钥文件必须放在 `.env` 中，代码需要主动加载 `.env`；不应进入版本控制的文件需要加入 `.gitignore`。

---

## 技术栈概览

### 后端 (Go)
- **Go 1.26.1+** - 主要后端语言
- **Connect RPC** (`connectrpc.com/connect`) - 类型安全的 RPC 框架
- **Protocol Buffers** + **buf** - API 定义与代码生成
- **Kong** (`github.com/alecthomas/kong`) - CLI 参数解析
- **Zerolog** (`github.com/rs/zerolog`) - 结构化日志
- **goutils** (`github.com/117503445/goutils`) - 通用工具库

### 前端 (React)
- **React 19** - UI 框架
- **TypeScript** - 类型安全
- **Vite 8** - 构建工具与开发服务器
- **Tailwind CSS 4** - 实用优先的 CSS 框架
- **Connect-ES** - TypeScript RPC 客户端
- **Lucide React** - 图标库
- **shadcn/ui 模式** - 组件架构

### 构建与开发工具
- **Task (go-task)** - 用于构建、运行、部署的任务运行器
- **Docker** - 容器化
- **GitHub Actions** - CI/CD
- **pnpm** - 前端包管理器
- **buf** - Protobuf 代码检查与生成

---

## 目录结构

```
project-root/
├── .env                        # 环境变量文件（已加入 gitignore）
├── .gitignore
├── buf.yaml                    # buf 代码检查/破坏性变更配置
├── buf.gen.yaml                # buf 代码生成配置
├── go.mod                      # Go 模块定义
├── go.sum
├── Taskfile.yml                # 主任务运行器配置
├── compose.yaml                # 开发环境 Docker Compose 配置
│
├── cmd/                        # 应用入口点
│   ├── cli/                    # CLI 应用
│   │   ├── main.go
│   │   ├── cli.go              # 带 kong 标签的 CLI 结构体
│   │   └── test.go
│   ├── rpc/                    # RPC 服务器
│   │   ├── main.go
│   │   ├── server.go
│   │   ├── handler.go
│   │   ├── context.go
│   │   └── static.go           # 静态文件服务
│   ├── fc-web/                 # 阿里云函数计算 (HTTP 触发)
│   │   ├── main.go
│   │   ├── server.go
│   │   ├── handler.go
│   │   └── context.go
│   ├── fc-event/               # 阿里云函数计算 (事件触发)
│   │   ├── main.go
│   │   └── context.go
│   └── fc-web-client/          # FC Web 客户端
│       └── main.go
│
├── internal/                   # 私有应用包
│   └── buildinfo/              # 构建时版本信息
│       └── info.go
│
├── pkg/                        # 公共包
│   └── rpc/                    # 生成的 RPC 代码
│       ├── template.proto      # Protobuf 定义文件
│       ├── template.pb.go      # 生成的 protobuf 代码
│       └── rpcconnect/
│           └── template.connect.go  # 生成的 Connect 代码
│
├── fe/                         # 前端 React 应用
│   ├── package.json
│   ├── pnpm-lock.yaml
│   ├── tsconfig.json
│   ├── tsconfig.app.json
│   ├── tsconfig.node.json
│   ├── vite.config.ts
│   ├── eslint.config.js
│   ├── index.html
│   ├── public/
│   ├── dist/                   # 构建输出
│   ├── node_modules/
│   └── src/
│       ├── main.tsx            # 入口文件
│       ├── App.tsx             # 根组件
│       ├── index.css           # 全局样式（含 Tailwind）
│       ├── assets/
│       ├── components/
│       │   └── ui/             # shadcn 风格组件
│       │       └── card.tsx
│       ├── gen/                # 生成的 RPC 客户端代码
│       └── lib/
│           ├── rpc.ts          # RPC 客户端配置
│           └── utils.ts        # 工具函数（cn 等）
│
├── scripts/                    # 构建与部署脚本
│   ├── entrypoint.sh           # Docker 入口脚本
│   ├── docker/
│   │   └── dev.Dockerfile
│   ├── go-scripts/             # Go 编写的构建脚本
│   │   ├── main.go
│   │   ├── cli.go
│   │   ├── build.go
│   │   ├── release.go
│   │   ├── deploy.go
│   │   ├── build-docker.go
│   │   ├── format.go
│   │   └── invoke.go
│   ├── tasks/                  # Taskfile 包含文件
│   │   ├── base/
│   │   ├── build/
│   │   ├── run/
│   │   ├── gen/
│   │   ├── deploy/
│   │   ├── format/
│   │   ├── fe/
│   │   └── test/
│   └── e2e/                    # E2E 测试（Python Playwright）
│       ├── main.py
│       ├── pyproject.toml
│       ├── uv.lock
│       ├── cases/
│       └── lib/
│
├── data/                       # 构建输出与运行时数据
│   ├── cli/
│   ├── rpc/
│   ├── fc-event/
│   ├── fc-web/
│   ├── e2e/
│   └── release/
│
├── docs/
└── .github/
    └── workflows/
        └── master.yml          # CI/CD 流水线
```

---

## Go 代码风格

### 日志模式

使用 zerolog 配合 goutils 封装实现结构化日志：

```go
import (
    "github.com/117503445/goutils/glog"
    "github.com/rs/zerolog/log"
)

// init 初始化结构化日志。
func init() {
    glog.InitZeroLog()
}

// main 输出结构化日志示例。
func main() {
    log.Info().
        Str("key", value).
        Interface("data", obj).
        Msg("操作描述")
}
```

### 使用 Kong 的 CLI 模式

使用 kong 标签定义 CLI 结构体：

```go
// cli.go
var cli struct {
    Test  struct{} `cmd:"" help:"运行测试"`
    Build struct {
        Output string `short:"o" help:"输出目录"`
    } `cmd:"" help:"构建应用"`
}

// main.go
// main 解析 CLI 命令并执行。
func main() {
    ctx := kong.Parse(&cli)
    log.Info().Interface("cli", cli).Send()
    if err := ctx.Run(); err != nil {
        log.Fatal().Err(err).Msg("运行失败")
    }
}
```

### RPC Handler 模式

```go
// handler.go
type TemplateHandler struct {
    // 依赖项
}

// NewTemplateHandler 创建模板 RPC 处理器。
func NewTemplateHandler() *TemplateHandler {
    return &TemplateHandler{}
}

// Healthz 返回服务健康状态和版本信息。
func (h *TemplateHandler) Healthz(
    ctx context.Context,
    req *connect.Request[rpc.HealthzRequest],
) (*connect.Response[rpc.ApiResponse], error) {
    resp := connect.NewResponse(&rpc.ApiResponse{
        Code:    0,
        Message: "success",
        Payload: &rpc.ApiResponse_Healthz{
            Healthz: &rpc.HealthzResponse{
                Version: buildinfo.GitVersion,
            },
        },
    })
    return resp, nil
}
```

### Buildinfo 模式

在构建时注入版本信息：

```go
// internal/buildinfo/info.go
package buildinfo

var (
    GitCommit  = ""
    GitBranch  = ""
    GitTag     = ""
    GitDirty   = ""
    GitVersion = ""
    BuildTime  = ""
    BuildDir   = ""
)

// 构建命令示例：
// go build -ldflags "-X github.com/117503445/go-template/internal/buildinfo.GitCommit=$(git rev-parse HEAD)"
```

---

## Protobuf API 模式

### Proto 文件结构

```protobuf
syntax = "proto3";
package pkg.rpc;
option go_package = "github.com/username/project/pkg/rpc;rpc";

// 统一响应包装器
message ApiResponse {
    int64 code = 1;
    string message = 2;
    oneof payload {
        HealthzResponse healthz = 3;
        // 在此添加更多响应类型
    }
}

message HealthzRequest {}
message HealthzResponse {
    string version = 1;
}

service TemplateService {
    rpc Healthz(HealthzRequest) returns (ApiResponse);
}
```

### buf.gen.yaml 配置

```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: .
    opt: paths=source_relative
  - local: protoc-gen-connect-go
    out: .
    opt: paths=source_relative
  - local: protoc-gen-es
    out: fe/src/gen
    opt: target=ts
  - local: protoc-gen-connect-es
    out: fe/src/gen
    opt: target=ts
```

---

## React/TypeScript 代码风格

### 组件模式

使用函数式组件 + Hooks：

```tsx
import { useEffect, useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

interface DataState {
  value: string
  status: 'loading' | 'success' | 'error'
}

// MyComponent 展示数据加载状态。
function MyComponent() {
  const [data, setData] = useState<DataState>({
    value: '',
    status: 'loading',
  })

  useEffect(() => {
    // 获取数据
  }, [])

  return (
    <Card>
      <CardHeader>
        <CardTitle>标题</CardTitle>
      </CardHeader>
      <CardContent>
        {/* 内容 */}
      </CardContent>
    </Card>
  )
}

export default MyComponent
```

### RPC 客户端配置

```typescript
// lib/rpc.ts
import { createConnectTransport } from '@connectrpc/connect-web'
import { TemplateService } from '@/gen/template_connect'

const transport = createConnectTransport({
  baseUrl: '',
})

export const rpcClient = new TemplateService(transport)
```

### Vite 配置

```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import path from 'path'

export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    proxy: {
      '/pkg.rpc': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

### 静态资源嵌入

前端生产构建输出到 `fe/dist` 后，需要同步到后端 package 内的静态资源目录，再通过 Go `embed` 嵌入后端。`//go:embed` 只能嵌入当前 package 目录下的文件，不能直接嵌入仓库根目录下的 `fe/dist`。构建任务必须保证先运行 `fe:build`，再同步静态资源并编译后端二进制文件。

```go
// cmd/rpc/static.go
package main

import (
    "embed"
    "io/fs"
    "net/http"
)

// staticFS 嵌入前端生产构建产物。
//
//go:embed all:static/dist
var staticFS embed.FS

// newStaticHandler 创建静态资源处理器。
func newStaticHandler() http.Handler {
    distFS, err := fs.Sub(staticFS, "static/dist")
    if err != nil {
        return http.NotFoundHandler()
    }
    return http.FileServer(http.FS(distFS))
}
```

### Tailwind CSS 配置

```css
/* index.css */
@import "tailwindcss";

/* 自定义主题变量 */
@theme {
  --color-primary-500: #3b82f6;
  --color-primary-600: #2563eb;
  --color-primary-700: #1d4ed8;
  --color-accent-500: #8b5cf6;
}
```

---

## 任务运行器 (go-task)

### Taskfile 最佳实践

- 根目录 `Taskfile.yml` 只做全局配置、`includes` 和稳定入口转发，不写具体业务逻辑。
- 领域任务放到 `scripts/tasks/<领域>/Taskfile.yml`，例如 `build`、`run`、`gen`、`deploy`、`format`、`fe`、`test`。
- 所有对用户可见的任务必须有中文 `desc`；复用但不希望直接调用的任务设置 `internal: true`。
- 跨 include 依赖使用完整命名空间，例如 `:base:clear`、`:build:bin`；依赖链不能依赖当前 shell 的隐式工作目录。
- 任务命名使用小写短横线或既有项目风格；同一仓库内保持一致。
- Taskfile 不能包含复杂逻辑，越简单越好；复杂逻辑写在 `scripts/go-scripts/` 中，再由 Taskfile 调用。
- 所有任务都要确保基于本地最新代码执行；运行、测试、E2E 等任务必须通过 `deps` 依赖必要的生成或构建任务，例如 E2E 依赖后端和前端构建。
- 编译、代码生成等无副作用任务应写好 `sources` 和 `generates`，确保代码不变时不重新执行。
- 有副作用的任务不要配置 `sources` 和 `generates`，例如 `run`、`deploy`、`test`、`e2e`。
- 需要传递用户参数时使用 `{{.CLI_ARGS}}`，例如 `go-task test:e2e -- --case <name>`。
- 前端、E2E 等子项目优先使用 `dir`，避免在命令里反复 `cd`。
- 需要临时补充工具链路径时，在任务内设置 `env`。
- 修改 Taskfile 后运行 `task --list` 检查任务是否可发现，必要时用 `task --dry <任务名>` 预览执行计划。
- 需要跳过缓存强制执行时使用 `task --force <任务名>`。

### Taskfile 修改流程

1. 先用 `rg -n "任务名|脚本名|文件名" Taskfile.yml scripts/tasks` 查清现有引用。
2. 判断是否需要新增 include；如果是新领域，创建 `scripts/tasks/<领域>/Taskfile.yml` 并在根 `Taskfile.yml` 注册。
3. 为每个用户可见任务补中文 `desc`。
4. 确认顶层只保留 `run` 和 `test` 稳定入口。
5. 为运行和测试任务补齐构建依赖，确保任务基于本地最新代码执行。
6. 为编译、代码生成等无副作用任务补齐 `sources` 和 `generates`。
7. 不给有副作用的任务配置 `sources` 和 `generates`。
8. 跨领域依赖写成 `:领域:任务`。
9. 修改后运行 `task --list`，再按项目要求运行 `go-task test`。

### 主 Taskfile.yml

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

### 测试任务示例

```yaml
# scripts/tasks/run/Taskfile.yml
version: 3

tasks:
  rpc:
    desc: "运行 RPC 服务"
    deps:
      - ":build:bin"
    cmds:
      - ./data/rpc/rpc {{.CLI_ARGS}}
```

```yaml
# scripts/tasks/test/Taskfile.yml
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
    dir: ./scripts/e2e
    cmds:
      - uv run main.py {{.CLI_ARGS}}
```

### 构建任务缓存示例

```yaml
# scripts/tasks/build/Taskfile.yml
version: 3

tasks:
  bin:
    desc: "构建后端二进制文件"
    deps:
      - ":gen:rpc"
      - ":fe:build"
    sources:
      - ./cmd/**
      - ./internal/**
      - ./pkg/**
      - ./cmd/rpc/static/dist/**
      - ./go.mod
      - ./go.sum
      - ./scripts/go-scripts/**
      - ./scripts/tasks/**
    generates:
      - ./data/cli/cli
      - ./data/rpc/rpc
    cmds:
      - rm -rf ./cmd/rpc/static/dist
      - cp -R ./fe/dist ./cmd/rpc/static/dist
      - go run ./scripts/go-scripts build
```

```yaml
# scripts/tasks/fe/Taskfile.yml
version: 3

tasks:
  build:
    desc: "构建前端"
    dir: ./fe
    sources:
      - ./package.json
      - ./pnpm-lock.yaml
      - ./index.html
      - ./vite.config.ts
      - ./src/**
    generates:
      - ./dist/**
    cmds:
      - pnpm build
```

### 集成测试脚本模式

`scripts/go-scripts it` 负责集成测试的复杂逻辑：启动已编译的 Go 服务，等待健康检查通过，使用 Go client 调用服务接口验证行为，并在结束时清理进程。Taskfile 只调用该命令并收集日志。

```go
// scripts/go-scripts/it.go
package main

// runIT 启动服务并通过 Go client 验证接口行为。
func runIT() error {
    // 启动 ./data/rpc/rpc，并把服务日志写入 ./data/it/server.log
    // 等待 /healthz 或 RPC healthz 就绪
    // 使用 Go client 调用服务并断言响应
    // 测试结束后清理服务进程
    return nil
}
```

### 常用任务命令

```bash
# 从 proto 生成 RPC 代码
task gen:gen-rpc

# 构建二进制文件
task build:bin

# 构建发布版本二进制文件
task build:release

# 构建并推送 Docker 镜像
task build:docker -- --push

# 运行 RPC 服务器
task run:rpc

# 运行前端开发服务器
task fe:dev

# 运行本地服务
go-task run

# 单独运行单元测试
go-task test:ut

# 单独运行集成测试
go-task test:it

# 单独调试指定 E2E 用例
go-task test:e2e -- --case <name>

# 运行所有测试（必须包含 E2E）
go-task test
```

---

## E2E 测试

使用 Python + uv + Playwright 编写 E2E 测试。项目有前端时，`go-task test` 必须运行完整的前后端服务，并使用 Python Playwright 编写浏览器调用代码。

每个 case 运行后，都要在 `./data/e2e/<case_name>/` 下输出日志、截图和测试报告；截图需要覆盖所有关键点。测试报告使用 Playwright 自定义 Markdown 报告器实现，参考 https://playwright.net.cn/docs/test-reporters，并在报告中引用关键截图。

```python
# scripts/e2e/main.py
def run_test(case_name: str, output_dir: Path, page: Page, logger: logging.Logger) -> bool:
    """运行单个 E2E 用例，并输出日志、截图和 Markdown 报告。"""
    case_dir = output_dir / case_name
    screenshots_dir = case_dir / "screenshots"
    logs_dir = case_dir / "logs"
    report_path = case_dir / "report.md"

    screenshots_dir.mkdir(parents=True, exist_ok=True)
    logs_dir.mkdir(parents=True, exist_ok=True)

    # 关键点截图需要写入报告，便于回放失败现场
    page.screenshot(path=screenshots_dir / "home.png", full_page=True)
    report_path.write_text("# E2E 测试报告\n\n![首页截图](screenshots/home.png)\n", encoding="utf-8")
    return True
```

运行方式：
```bash
# 先运行新增或变更的单个 E2E 用例
go-task test:e2e -- --case <name>

# 单个用例通过后运行全部测试
go-task test
```

---

## GitHub Actions CI/CD

```yaml
# .github/workflows/master.yml
on:
  push:
    branches: [master]

jobs:
  master:
    runs-on: ubuntu-latest
    container:
      image: your-dev-image
    steps:
      - name: ci
        run: |
          git clone https://github.com/  ${{ github.repository }}.git /workspace
          cd /workspace
          git checkout ${{ github.sha }}
          go-task build:docker -- --push
          go-task build:release
          # 使用 gh CLI 创建 Release
```

---

## 项目初始化检查清单

初始化新项目时，请完成以下步骤：

1. **项目基础设置**
   - [ ] 初始化 Go 模块：`go mod init github.com/username/project`
   - [ ] 创建目录结构：`cmd/`、`internal/`、`pkg/`、`scripts/`、`fe/`、`data/`
   - [ ] 设置 `.gitignore` 和 `.env.example`

2. **后端设置**
   - [ ] 创建 `buf.yaml` 和 `buf.gen.yaml`
   - [ ] 在 `pkg/rpc/` 中定义 proto 文件
   - [ ] 在 `cmd/` 中创建主入口文件
   - [ ] 设置 buildinfo 包
   - [ ] 在 `scripts/go-scripts/` 中创建构建脚本

3. **前端设置**
   - [ ] 使用 Vite 初始化：`pnpm create vite`
   - [ ] 添加依赖：React、Tailwind CSS、Connect-ES
   - [ ] 配置 `vite.config.ts` 路径别名
   - [ ] 设置 ESLint 配置
   - [ ] 创建组件结构
   - [ ] 配置 `fe:build` 输出 `fe/dist`

4. **开发工具配置**
   - [ ] 创建 `Taskfile.yml` 并包含子任务文件
   - [ ] 配置顶层 `run` 入口，且 `go-task run` 必须基于最新构建产物运行
   - [ ] 配置顶层 `test` 入口，不配置顶层 `e2e` 入口
   - [ ] 配置 `test:ut`、`test:it` 和 `test:e2e`，且 `go-task test` 必须包含三类测试
   - [ ] 设置 Docker 配置
   - [ ] 配置 GitHub Actions

5. **代码生成**
   - [ ] 运行 `buf generate` 生成 RPC 代码
   - [ ] 验证 Go 和 TypeScript 代码生成结果
   - [ ] 确认前端 `fe/dist` 已通过 Go `embed` 嵌入后端

---

## 关键模式遵循指南

1. **cmd/internal/pkg 分离**：入口点在 `cmd/`，私有代码在 `internal/`，公共包在 `pkg/`

2. **统一 API 响应**：使用包含 `code`、`message` 和 `oneof payload` 的包装响应

3. **结构化日志**：始终使用带结构化字段的 zerolog

4. **Kong CLI**：将 CLI 定义为带 kong 标签的结构体，使用 `ctx.Run()` 模式

5. **构建时信息**：通过 ldflags 在构建时注入版本信息

6. **路径别名**：前端使用 `@/` 作为导入路径前缀

7. **前端嵌入后端**：前端需要先编译为 `fe/dist`，再通过 Go `embed` 嵌入到后端服务中；生产运行时由后端直接服务静态资源

8. **组件架构**：React 组件遵循 shadcn/ui 模式

9. **任务化构建**：所有构建、运行、部署操作均使用 go-task 管理
