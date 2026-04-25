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
│   │   └── e2e/
│   └── e2e/                    # E2E 测试（Python/playwright）
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

func init() {
    glog.InitZeroLog()
}

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

func NewTemplateHandler() *TemplateHandler {
    return &TemplateHandler{}
}

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
  e2e:
    taskfile: ./scripts/tasks/e2e
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

# 运行 E2E 测试
task e2e:test
```

---

## E2E 测试

使用 Python + uv 运行 E2E 测试：

```python
# scripts/e2e/main.py
def run_test(output_dir: Path, screenshots_dir: Path, logs_dir: Path, logger: logging.Logger) -> bool:
    # 测试实现
    return True
```

运行方式：
```bash
cd scripts/e2e && uv run main.py
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

4. **开发工具配置**
   - [ ] 创建 `Taskfile.yml` 并包含子任务文件
   - [ ] 设置 Docker 配置
   - [ ] 配置 GitHub Actions

5. **代码生成**
   - [ ] 运行 `buf generate` 生成 RPC 代码
   - [ ] 验证 Go 和 TypeScript 代码生成结果

---

## 关键模式遵循指南

1. **cmd/internal/pkg 分离**：入口点在 `cmd/`，私有代码在 `internal/`，公共包在 `pkg/`

2. **统一 API 响应**：使用包含 `code`、`message` 和 `oneof payload` 的包装响应

3. **结构化日志**：始终使用带结构化字段的 zerolog

4. **Kong CLI**：将 CLI 定义为带 kong 标签的结构体，使用 `ctx.Run()` 模式

5. **构建时信息**：通过 ldflags 在构建时注入版本信息

6. **路径别名**：前端使用 `@/` 作为导入路径前缀

7. **组件架构**：React 组件遵循 shadcn/ui 模式

8. **任务化构建**：所有构建、运行、部署操作均使用 go-task 管理