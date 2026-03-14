# vibe-init 嵌入 CLAUDE.md 和 skills 目录

## 主要内容和目的

将 `scripts/vibe-init/CLAUDE.md` 和 `scripts/vibe-init/skills/` 目录嵌入到 Go 程序中，运行时自动释放到 `~/.claude/` 目录。

## 更改内容描述

### 修改文件

- `scripts/vibe-init/main.go`
  - 扩展 `//go:embed` 指令，嵌入 `settings.json`、`CLAUDE.md` 和整个 `skills` 目录
  - 新增 `copyFS()` 辅助函数，递归复制嵌入的目录到文件系统
  - 运行时释放 `CLAUDE.md` 到 `~/.claude/CLAUDE.md`
  - 运行时释放 `skills/` 目录到 `~/.claude/skills/`

- `scripts/vibe-init/settings.json`
  - 更新 Claude 配置，添加 hooks、plugins 等配置项

### 新增文件

- `scripts/vibe-init/CLAUDE.md` - Claude 项目级配置
- `scripts/vibe-init/skills/commit/SKILL.md` - commit skill 定义
- `scripts/vibe-init/skills/init-project/SKILL.md` - init-project skill 定义

## 验证方法和结果

```bash
cd /workspace/project/docker-dev/scripts/vibe-init && go build -o /dev/null .
```

构建成功，无错误输出。