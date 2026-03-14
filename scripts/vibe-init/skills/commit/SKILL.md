---
name: commit
description: 提交 commit 的必做事项
---

# commit

1. stage 所有变更文件

```sh
git add -A
```

2. 查看当前变更文件，总结出 COMMIT_NAME(中文)，要求简洁明了，能够概括本次变更的主要内容。

3. 移动 `prompts.jsonl` 和 `results.jsonl` 文件

```sh
DATE=$(date +%Y%m%d.%H%M%S)

DIR_COMMIT=docs/commits/$DATE.$COMMIT_NAME

mkdir -p docs/commits

if [ -f prompts.jsonl ]; then
    mv prompts.jsonl $DIR_COMMIT/prompts.jsonl
fi

if [ -f responses.jsonl ]; then
    mv responses.jsonl $DIR_COMMIT/responses.jsonl
fi

touch $DIR_COMMIT/README.md
```

4. 检查 `prompts.jsonl` 和 `responses.jsonl` 文件内容，如果有敏感信息和密钥，务必删除敏感信息。

5. 检查是否有二进制可执行程序、构建缓存等不适合加入版本控制系统的文件。如果有，加入 gitignore 并取消 stage。

6. 根据当前变更内容，在 `$DIR_COMMIT/README.md` 编写中文说明，包含以下内容：

- 主要内容和目的
- 更改内容描述
- 验证方法和结果

7. 再次 stage 变更文件，并使用中文提交 commit。