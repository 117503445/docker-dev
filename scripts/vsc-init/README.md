# vsc-init

Code Server 安装扩展、修改设置

## 使用方法

安装

```sh
go install github.com/117503445/vsc-init@latest
```

运行

```sh
vsc-init

EXTS=golang.go,njzy.stats-bar vsc-init # 使用 EXTS 环境变量安装额外的扩展
```

`vsc-init` 会根据 `pkg/assets/assets.go`，对于本地已安装的 Code Server，进行

- 基于本地拓展版本，计算出需要安装/升级的拓展列表，并下载安装
- 写入 Settings 配置
- 写入 Keybindings 配置

## 注意事项

- VSCode 中常使用 Pylance 提供 Python 语言支持，但是在 Code Server 等非官方 VSCode 环境中无法使用 Pylance。即使强行安装上，Pylance 也会检测到处于非官方环境，从而拒绝启动。所以使用 `detachhead.basedpyright` 作为替代。
- 安装了 [vscode-key-runner](https://github.com/117503445/vscode-key-runner) 并完成相关配置，允许按下 F5 在终端调用 `go-task` 命令。
