# 添加 sshole agent 支持

## 主要内容和目的

为 dev-desktop 镜像添加 sshole agent 自动启动功能，当设置 `SSHOLE_AGENT_HUB_SERVER` 环境变量时自动启动 sshole agent。

## 更改内容描述

1. **Dockerfile** (`images/dev-desktop/Dockerfile`)
   - 添加 `go install github.com/117503445/sshole/cmd/agent@latest` 安装 sshole agent

2. **entrypoint** (`scripts/entrypoint/main.go`)
   - 新增 sshole agent 启动逻辑
   - 当 `SSHOLE_AGENT_HUB_SERVER` 不为空时启动 agent
   - 支持环境变量配置：`SSHOLE_AGENT_AUTH`、`SSHOLE_AGENT_NAME`、`SSHOLE_AGENT_LOCAL_PORT`、`SSHOLE_AGENT_SKIP_SSHD`
   - 日志输出到 `/docker-dev/logs/sshole-agent.log`

3. **README.md**
   - 在"运行时可配置环境变量"表格中添加 sshole agent 相关环境变量说明

## 验证方法和结果

- entrypoint 代码编译通过：`go build -o /dev/null .` 无错误
- sshole agent 参数格式通过 `agent --help` 确认