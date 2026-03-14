# Dockerfile 多阶段构建重构

## 主要内容和目的

将现有的 Docker 构建流程重构为多阶段构建模式，将 entrypoint、vsc-init、shell-init、vibe-init 分别独立构建，并在主 Dockerfile 最后统一复制和运行。

## 更改内容描述

### 新增文件

| 文件 | 描述 |
|------|------|
| `scripts/shell-init/go.mod` | Go module 定义 |
| `scripts/shell-init/go.sum` | Go 依赖校验 |
| `scripts/shell-init/main.go` | Shell 初始化逻辑（zsh、oh-my-zsh、histmon） |
| `scripts/vibe-init/go.mod` | Go module 定义 |
| `scripts/vibe-init/go.sum` | Go 依赖校验 |
| `scripts/vibe-init/main.go` | Claude Code 设置逻辑 |

### 修改文件

| 文件 | 变更 |
|------|------|
| `images/dev-desktop/Dockerfile` | 重构为多阶段构建 |
| `scripts/vsc-init/pkg/assets/assets.go` | 添加 `GitHub.copilot-chat` 扩展 |

### 多阶段构建结构

1. **构建阶段** (使用 `golang:1.24.1`):
   - `builder-entrypoint` - 构建 entrypoint 二进制
   - `builder-vsc-init` - 构建 vsc-init 二进制
   - `builder-shell-init` - 构建 shell-init 二进制
   - `builder-vibe-init` - 构建 vibe-init 二进制

2. **主阶段** (使用 `lscr.io/linuxserver/webtop:arch-kde`):
   - 从构建阶段复制所有二进制文件
   - 通过 `go install` 安装 histmon（保持原有方式）
   - 复制资源文件（`.zshrc`, `settings.json`）
   - 按顺序运行初始化脚本: `vsc-init` → `shell-init` → `vibe-init`

3. **CHINA_MIRROR 支持**: 使用 `--build-arg GO_BASE_IMAGE=registry.cn-hangzhou.aliyuncs.com/117503445-mirror/sync:linux.amd64.docker.io.library.golang.1.24.1` 加速国内构建

## 验证方法和结果

```bash
go-task build-dev-desktop
go-task test
```

测试结果：38 项测试全部通过

```
Go:                  ✓ PASS
Rust/Cargo:          ✓ PASS
Python:              ✓ PASS
...
Claude Code:         ✓ PASS

Total: 38 passed, 0 failed
```