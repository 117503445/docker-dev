# 添加 histmon 工具

## 主要内容和目的

将 histmon（命令执行监控工具）集成到项目中，作为 scripts 目录下的一个子项目。

## 更改内容描述

1. **新增 scripts/histmon 目录**
   - 从 github.com/117503445/histmon 复制源代码
   - 修改 go.mod 模块路径为 `github.com/117503445/docker-dev/scripts/histmon`
   - 移除不必要的文件（.git、.gitignore、compose.yaml、dev.Dockerfile、Taskfile.yml）

2. **修改 images/dev-desktop/Dockerfile**
   - 添加 `builder-histmon` 构建阶段
   - 移除 `go install github.com/117503445/histmon@master` 远程安装
   - 改为从本地源码构建并复制到 `/usr/local/bin/histmon`

3. **更新 README.md**
   - 在 `go install` 工具列表中添加 histmon 工具的安装说明

## 验证方法和结果

- 检查 scripts/histmon 目录结构正确
- Dockerfile 构建阶段语法正确
- README.md 文档更新完整