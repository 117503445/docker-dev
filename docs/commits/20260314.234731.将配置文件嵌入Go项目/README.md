# 将配置文件嵌入 Go 项目

## 主要内容和目的

将 `.zshrc` 和 Claude `settings.json` 配置文件从 `assets` 目录移动到各自的 Go 项目中，使用 Go 的 `embed.FS` 功能嵌入到二进制文件中，简化 Dockerfile 配置。

## 更改内容描述

### 文件移动

- `assets/.zshrc` → `scripts/shell-init/.zshrc`
- `assets/claude/settings.json` → `scripts/vibe-init/settings.json`

### Go 代码修改

**scripts/shell-init/main.go:**
- 使用 `//go:embed .zshrc` 嵌入配置文件
- 移除从 `/tmp/.zshrc` 读取的逻辑
- 直接写入嵌入的 `.zshrc` 内容到 `$HOME/.zshrc`
- 自动创建 `.zshrc-custom` 文件

**scripts/vibe-init/main.go:**
- 使用 `//go:embed settings.json` 嵌入配置文件
- 移除从 `/tmp/claude-settings.json` 读取的逻辑
- Claude settings 更新为 `bypassPermissions` 模式

### Dockerfile 修改

**images/dev/Dockerfile:**
- 移除 `COPY ./assets/.zshrc /root/.zshrc`
- 移除 `RUN touch /root/.zshrc-custom`
- 移除 `RUN TOKEN="T" ENDPOINT="E" histmon install`（已集成到 shell-init）
- 添加编译和运行 shell-init/vibe-init 的步骤

**images/dev-desktop/Dockerfile:**
- 移除 `COPY ./assets/.zshrc /tmp/.zshrc`
- 移除 `COPY ./assets/claude/settings.json /tmp/claude-settings.json`
- 已正确配置使用多阶段构建

### .zshrc 新增内容

新增别名和函数：
```bash
alias cla="IS_SANDBOX=1 claude --dangerously-skip-permissions"
alias cod="codex --dangerously-bypass-approvals-and-sandbox"
yay-su() { su - builder -c "yay -Su $* --noconfirm"; }
yay-syu() { su - builder -c "yay -Syu $* --noconfirm"; }
alias pacman-sy="pacman -Sy --noconfirm"
alias pacman-syu="pacman -Syu --noconfirm"
```

## 验证方法和结果

```bash
# 编译验证
cd scripts/shell-init && go build . && echo "OK"
cd scripts/vibe-init && go build . && echo "OK"
```

- shell-init 编译成功
- vibe-init 编译成功
- embed 功能正常工作