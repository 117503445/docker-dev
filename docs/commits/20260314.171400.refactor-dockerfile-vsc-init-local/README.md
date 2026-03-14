# Dockerfile多阶段构建及vsc-init本地化

## 主要内容和目的

1. 改造 `images/dev-desktop/Dockerfile`，使用多阶段构建优化镜像
2. 将 vsc-init 代码复制到本仓库 `scripts/vsc-init/` 目录，便于本地引用
3. 添加 anthropic.claude-code VS Code 扩展

## 更改内容描述

### Dockerfile 改造
- 改为在镜像内直接编译 Go 二进制文件（entrypoint 和 vsc-init）
- 避免从 Docker Hub 拉取额外的 golang 镜像
- 对网络依赖的操作添加容错处理（`|| echo "Warning: ..."`）

### vsc-init 本地化
- 将 vsc-init 源码复制到 `scripts/vsc-init/`
- Dockerfile 从本地复制并编译 vsc-init
- 修复 `getVscodeEngine()` 函数，使用正则表达式更健壮地解析 code-server 版本

### 扩展添加
- 在 `scripts/vsc-init/pkg/assets/assets.go` 中添加 `anthropic.claude-code` 扩展

### Taskfile 更新
- 移除 `build-dev-desktop` 任务中的预编译步骤（现在由 Dockerfile 内部处理）

## 验证方法和结果

运行 `go-task test` 验证容器环境：

```
========================================
Total: 38 passed, 0 failed
========================================
All tests passed!
```

所有 38 项测试通过，包括：
- 编程语言：Go, Rust, Python, Node.js, Java, .NET, GCC, Clang, CMake
- DevOps 工具：Docker, Docker Compose, Podman, kubectl, Helm
- 工具：Git, GitHub CLI, Task, Make, uv, Typst
- 编辑器：Vim, Nano, Micro
- Claude Code