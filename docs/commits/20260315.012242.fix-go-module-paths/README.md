# 修正 Go 模块路径

## 主要内容和目的

修正 scripts 目录下 Go 项目的模块路径，使其可以通过 `go install` 正确安装，并在 README.md 中添加安装说明。

## 更改内容描述

1. **修正 go.mod 模块路径**
   - `scripts/entrypoint/go.mod`: `github.com/117503445/docker-dev/entrypoint` → `github.com/117503445/docker-dev/scripts/entrypoint`
   - `scripts/shell-init/go.mod`: `github.com/117503445/docker-dev/shell-init` → `github.com/117503445/docker-dev/scripts/shell-init`
   - `scripts/vibe-init/go.mod`: `github.com/117503445/docker-dev/vibe-init` → `github.com/117503445/docker-dev/scripts/vibe-init`
   - `scripts/vsc-init/go.mod`: `github.com/117503445/vsc-init` → `github.com/117503445/docker-dev/scripts/vsc-init`

2. **更新 vsc-init 内部导入**
   - `main.go` 和 `pkg/ext/ext.go` 中的导入路径已更新为新模块路径

3. **添加 README.md 安装说明**
   - 在 README.md 中添加"通过 go install 安装 scripts 下的工具"章节

4. **更新 .gitignore**
   - 添加编译生成的二进制文件到 .gitignore

## 验证方法和结果

```sh
# 验证构建
cd scripts/vsc-init && go build .
cd ../shell-init && go build .
cd ../vibe-init && go build .
cd ../entrypoint && go build .
```

所有模块构建成功，无错误。