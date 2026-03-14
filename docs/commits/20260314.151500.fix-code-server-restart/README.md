# 修复 code-server 不断重启问题

## 主要内容和目的

解决 Docker 容器中 code-server 服务不断重启的问题，使容器能够稳定运行。

## 更改内容描述

### 1. Taskfile.yml
- 将默认任务从 `build-dev` 改为 `run-dev-desktop`

### 2. scripts/entrypoint/main.go

| 问题 | 修复方案 |
|------|----------|
| 配置文件路径不匹配 | 添加 `--config` 参数显式指定配置文件路径 |
| 配置目录不存在时写入失败 | 添加目录创建逻辑 |
| `goutils.WriteText` 写入失败 | 改用 `os.WriteFile` 确保文件权限正确 |
| sshd 后台启动导致 s6 监管 deadlock | 添加 `-D` 参数使 sshd 前台运行 |
| s6 重启时端口冲突 | 添加端口占用检测逻辑，避免重复启动 |

## 验证方法和结果

1. 运行 `task run-dev-desktop` 启动容器
2. 访问 `http://localhost:44444` 进入 code-server
3. 使用密码 `123456` 登录
4. 观察容器稳定运行，code-server 和 sshd 不再重启

**验证结果**: 容器稳定运行，code-server 和 sshd 服务正常工作。