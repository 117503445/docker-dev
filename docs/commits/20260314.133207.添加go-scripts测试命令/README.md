# 添加 go-scripts 测试命令

## 主要内容和目的

在 `/workspace/project/docker-dev/scripts/go-scripts/` 目录下创建 Go 项目，添加 `test` 命令用于在 `task build-dev-desktop` 后验证容器环境是否正确配置。

## 更改内容描述

### 新增文件
- `scripts/go-scripts/go.mod` - Go 模块定义
- `scripts/go-scripts/go.sum` - 依赖锁定文件
- `scripts/go-scripts/main.go` - 入口文件，使用 kong CLI 框架
- `scripts/go-scripts/cli.go` - CLI 命令定义
- `scripts/go-scripts/test.go` - test 命令实现

### 修改文件
- `Taskfile.yml` - 添加 `test` 任务

### 测试覆盖的工具
- 编程语言：Go, Rust/Cargo, Python, Node.js, Java, .NET, GCC, Clang, CMake
- DevOps 工具：Docker, Docker Compose, Podman, kubectl, Helm
- 版本控制：Git, GitHub CLI
- 构建工具：Task, Make
- 其他：uv, Typst, gopls, golangci-lint, 编辑器 (Vim, Nano, Micro), Claude Code

## 验证方法和结果

```bash
# 编译验证
cd scripts/go-scripts && go build -o .

# 运行测试（需要 dev-desktop 容器运行中）
task test
```

输出示例：
```
========================================
       Container Environment Test
========================================

Go:                  ✓ PASS
    go version go1.25.5 linux/amd64
Rust/Cargo:          ✓ PASS
    cargo 1.84.1
...
========================================
Total: 30 passed, 0 failed
========================================
```