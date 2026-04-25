# vibe-init

`vibe-init` 用于初始化本机 Claude Code 和 Codex 的全局配置，让两种编码代理共享同一套基础说明、skills 和会话日志 hook。

## 安装

可以直接通过 `go install` 安装最新版本：

```bash
go install github.com/117503445/docker-dev/scripts/vibe-init@latest
```

安装完成后执行：

```bash
vibe-init
```

## 初始化内容

执行后会写入以下路径：

| 目标 | 路径 | 作用 |
| --- | --- | --- |
| 公共日志脚本 | `~/.ai/vibe_hook.js` | 将 hook stdin 原样追加到 JSONL 日志 |
| Claude Code 设置 | `~/.claude/settings.json` | 配置模型环境变量、插件和 hooks |
| Claude Code 指令 | `~/.claude/AGENTS.md`、`~/.claude/CLAUDE.md` | `CLAUDE.md` 通过 `@~/.claude/AGENTS.md` 导入公共指令 |
| Claude Code skills | `~/.claude/skills/` | 安装本项目内置 skills |
| Codex 设置 | `~/.codex/config.toml`、`~/.codex/hooks.json` | 启用 `codex_hooks` 并配置 hooks |
| Codex 指令 | `~/.codex/AGENTS.md` | 安装 Codex 全局指令 |
| Codex skills | `~/.codex/skills/` | 安装本项目内置 skills |

## 会话日志

Claude Code 和 Codex 都会在用户提交 prompt 时记录 `request`，在代理停止响应时记录 `response`。hook 不解析、不包装 stdin 中的 JSON，只把原始输入原样追加写入。日志路径格式：

```text
~/.ai/<URL encode dir path>/codex.jsonl
~/.ai/<URL encode dir path>/claude.jsonl
```

例如工作目录 `/workspace/project/demo` 会写入：

```text
~/.ai/%2Fworkspace%2Fproject%2Fdemo/codex.jsonl
~/.ai/%2Fworkspace%2Fproject%2Fdemo/claude.jsonl
```

每一行都是 Claude Code 或 Codex 传给 hook 的原始 JSON 对象。

## Claude Code 配置依据

Claude Code 的官方文档说明：

- `settings.json` 是 Claude Code 的分层配置机制，用户级设置位于 `~/.claude/settings.json`，项目级设置位于 `.claude/settings.json` 和 `.claude/settings.local.json`。见 [Claude Code settings](https://code.claude.com/docs/en/settings)。
- `CLAUDE.md` 是 Claude Code 启动时读取的持久指令文件，用户级路径是 `~/.claude/CLAUDE.md`，项目级路径是 `./CLAUDE.md` 或 `./.claude/CLAUDE.md`。见 [How Claude remembers your project](https://code.claude.com/docs/en/memory)。
- Claude Code 不直接读取 `AGENTS.md`；如果仓库已有 `AGENTS.md`，官方建议创建 `CLAUDE.md` 并用 `@AGENTS.md` 导入它。本项目按这个模式写入 `~/.claude/CLAUDE.md`。见 [AGENTS.md section](https://code.claude.com/docs/en/memory#agentsmd)。
- hooks 定义在 JSON settings 中，可放在 `~/.claude/settings.json`、`.claude/settings.json`、`.claude/settings.local.json` 等位置；`UserPromptSubmit`、`Stop`、`SessionStart`、`SessionEnd` 等是支持的事件。见 [Claude Code hooks reference](https://code.claude.com/docs/en/hooks)。
- skills 放在 `~/.claude/skills/` 或项目 `.claude/skills/` 下会被发现；每个 skill 是包含 `SKILL.md` 的目录。见 [Extend Claude with skills](https://code.claude.com/docs/en/skills)。

## Codex 配置依据

Codex 的官方文档说明：

- 用户级配置位于 `~/.codex/config.toml`，项目级覆盖可放在 `.codex/config.toml`。本项目只确保 `features.codex_hooks = true` 存在。见 [Codex Configuration Reference](https://developers.openai.com/codex/config-reference)。
- Codex 会在启动时读取 `AGENTS.md`。全局范围位于 Codex home，默认是 `~/.codex/AGENTS.md`；项目范围会从项目根目录走到当前目录读取 `AGENTS.override.md`、`AGENTS.md` 或 fallback 文件。见 [Custom instructions with AGENTS.md](https://developers.openai.com/codex/guides/agents-md)。
- Codex hooks 需要在 `config.toml` 中启用 `[features] codex_hooks = true`；hook 可以写在 `~/.codex/hooks.json` 或 `~/.codex/config.toml`，也支持项目级 `.codex/hooks.json` 或 `.codex/config.toml`。见 [Codex Hooks](https://developers.openai.com/codex/hooks)。
- Codex 的 `UserPromptSubmit` hook 输入包含 `prompt`，`Stop` hook 输入包含 `last_assistant_message`；本项目不解析这些字段，只按 `request`/`response` 类型保存原始 hook 输入。见 [Codex Hooks](https://developers.openai.com/codex/hooks)。
- Codex skills 可用于 CLI、IDE extension 和 Codex app；skill 是包含 `SKILL.md` 的目录，Codex 会按 metadata 发现并在需要时加载。见 [Codex Agent Skills](https://developers.openai.com/codex/skills)。

## 开发

运行测试：

```bash
task test
```
