---
name: commit
description: Git Commit 标准化操作流程
---

# 🚀 Git Commit 标准化流程

在提交代码前，请务必按以下步骤执行，确保代码库的整洁性、安全性及可追溯性。

### 1. 暂存变更 (Stage Changes)

首先将当前工作区的所有变更加入暂存区：

```bash
git add -A
```

不要只暂存自己变更的内容，要将所有变更内容暂存。

### 2. 安全检查与清理 (Security & Cleanup) ⚠️ **关键步骤**

仔细审查 `git status` 输出的文件列表，执行以下操作：

- **敏感信息排查**：严禁包含密钥（API Key）、密码、Token 或证书文件。若发现，提交到 `.gitignore`。

- **构建产物过滤**：检查是否意外包含了编译生成的二进制文件、日志文件或临时缓存。如有，请将其加入 `.gitignore`。

- **确认无误**：再次运行 `git status` 确认只保留了预期的源代码变更。

### 3. 编写提交信息 (Commit Message)

使用中文撰写清晰、规范的 Commit Message。

- **格式要求**：遵循 `类型(范围): 描述` 结构（如 `feat(auth): 新增登录接口`）。
- **内容要求**：
  - 简述做了什么（What）和为什么做（Why）。
  - 避免使用“修复 bug"、“更新代码”等模糊描述。

### 4. 执行提交 (Execute Commit)

```bash
git commit -m "你的提交信息"
```
