# 简化 Claude 权限配置为 bypassPermissions 模式

## 主要内容和目的

将 Claude Code 的权限配置从复杂的 allow 规则简化为 `bypassPermissions` 模式，适用于 Docker 容器等隔离开发环境。

## 更改内容描述

### 文件变更

- `assets/claude/settings.json`: 简化权限配置
  - 移除 `enabledPlugins` 配置（gopls-lsp 插件）
  - 移除复杂的 `permissions.allow` 规则列表
  - 添加 `defaultMode: "bypassPermissions"` 配置
  - 保留 `skipDangerousModePermissionPrompt: true`

### 配置对比

**修改前**（复杂的权限规则）:
```json
{
    "skipDangerousModePermissionPrompt": true,
    "enabledPlugins": { "gopls-lsp@claude-plugins-official": true },
    "permissions": {
        "allow": ["Read(**)", "Edit(**)", "Bash(*)", "LS(**)", "Grep(**)", "Glob(**)", "WebFetch", "MCP"],
        "deny": []
    }
}
```

**修改后**（简化配置）:
```json
{
    "skipDangerousModePermissionPrompt": true,
    "defaultMode": "bypassPermissions"
}
```

### 技术说明

原配置存在问题：
- `Read(**)` 等语法不规范，正确写法应为 `Read`（无括号）
- `LS(**)` 不是有效的工具名称
- `Grep`/`Glob` 无需单独配置，`Read` 规则会自动应用

`bypassPermissions` 模式跳过所有权限检查，适合在 Docker 容器等隔离环境中使用。

## 验证方法和结果

- 参考官方文档: https://code.claude.com/docs/en/permissions
- 确认配置语法正确
- 配置文件格式验证通过