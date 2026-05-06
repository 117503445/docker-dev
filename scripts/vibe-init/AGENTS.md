- 文档、注释 使用中文。每个方法和方法参数需要有注释，代码块内在必要的地方添加注释。写文档尽量保持精简。
- 如果通过 `go-task test` 命令可以运行测试，且需要修改代码，则
    1. 在测试用例中覆盖新增的需求。运行 `go-task test -- --case <name>`，预期会因为新增用例失败。
    2. 修改代码，运行 `go-task test -- --case <name>`，直到成功。
    3. 成功后，运行 `go-task test` 运行全部测试。密钥文件必须放在 .env 中，并且在代码中主动加载 .env。不应该放在 vcs 的文件（如代码生成的）应使用 .gitignore 排除。
- 前端使用 TypeScript 编写。如果没有特殊声明，其他所有代码、脚本使用 Go 编写。
- 如果需要 playwright，使用 https://github.com/playwright-community/playwright-go；如果需要写 go websocket，使用 https://github.com/coder/websocket。
- 如果使用 zerolog 作为日志库，使用 `log.Ctx(ctx)` 打印日志。确保通过日志可以快速定位问题。
- 当单个代码文件过大时，尝试拆分为多个文件或者多个模块。
- 如果存在 docs/需求.md，就在 docs/需求.md 记录来自用户的完整需求，不要遗漏，保持简洁。按逻辑结构组织需求，不按用户提出顺序追加。
