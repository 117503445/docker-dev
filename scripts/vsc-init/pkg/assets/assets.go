package assets

var Exts = []string{
	// 语言
	"jnoortheen.nix-ide",
	"golang.go",
	"ms-python.python",
	"ms-python.black-formatter",
	"detachhead.basedpyright",
	"fwcd.kotlin",
  "rust-lang.rust-analyzer",
	"tamasfe.even-better-toml",
	"bodil.prettier-toml",
	"zxh404.vscode-proto3",
	"redhat.vscode-xml",
	"redhat.vscode-yaml",
	"nvarner.typst-lsp",
	// Git
	"mhutchie.git-graph", // 清晰的 Git 可视化，操作便捷
	// IDE
	"ms-ceintl.vscode-language-pack-zh-hans",
	"pkief.material-icon-theme",
	"njzy.stats-bar", // 显示 CPU、内存、网络、磁盘占用
	"mechatroner.rainbow-csv",
  // Other
	// "alibaba-cloud.tongyi-lingma",
	"ms-azuretools.vscode-docker",
	"humao.rest-client", // 替代 Postman
	"tomoki1207.pdf",
  "iliazeus.vscode-ansi", // 显示 ANSI 文本颜色
	"anthropic.claude-code", // Claude Code VS Code extension
	"GitHub.copilot-chat",
}

var KeyBindings = `[
    {
        "key": "ctrl+alt+left",
        "command": "workbench.action.navigateBack",
        "when": "canNavigateBack"
    },
    {
        "key": "alt+left",
        "command": "-workbench.action.navigateBack",
        "when": "canNavigateBack"
    },
    {
        "key": "ctrl+alt+right",
        "command": "workbench.action.navigateForward",
        "when": "canNavigateForward"
    },
    {
        "key": "alt+right",
        "command": "-workbench.action.navigateForward",
        "when": "canNavigateForward"
    },
    {
        "key": "f5",
        "command": "key-runner.run"
    },
]`

var Settings = `{
  "editor.minimap.enabled": false,
  "editor.unicodeHighlight.allowedCharacters": {
    "！": true,
    "：": true,
    "，": true
  },
  "editor.wordWrap": "on",
  "explorer.confirmDelete": false,
  "explorer.confirmDragAndDrop": false,
  "editor.wordSeparators": "~!@#$%^&*()-=+[{]}\\|;:'\",.<>/！（）｛｝【】、；：’”，。《》？",
  "git.autofetch": true,
  "git.confirmSync": false,
  "git.ignoreMissingGitWarning": true,
  "git.enableSmartCommit": true,
  "go.testFlags": [
    "-v",
    "-count=1"
  ],
  "go.testTimeout": "24h",
  "go.toolsManagement.autoUpdate": true,
  "terminal.integrated.copyOnSelection": true,
  "terminal.integrated.fontFamily": "UbuntuMono Nerd Font Mono,UbuntuMono NF,Microsoft YaHei Mono",
  "terminal.integrated.fontWeightBold": "normal",
  "terminal.integrated.defaultProfile.linux": "zsh",
  "terminal.integrated.profiles.linux": {
    "zsh": {
      "path": "zsh"
    }
  },
  // "window.title": "${rootName}",
  //WINDOW
  "window.commandCenter": false,
  "workbench.layoutControl.enabled": false,
  "git.alwaysSignOff": true,
  "editor.inlineSuggest.enabled": true,
  "git.defaultBranchName": "master",
  "explorer.confirmPasteNative": false,
  "window.autoDetectColorScheme": true,
  "files.autoSave": "afterDelay",
  "workbench.iconTheme": "material-icon-theme",
  "redhat.telemetry.enabled": false,
  "terminal.integrated.commandsToSkipShell": [
    "key-runner.run"
  ],
  "terminal.integrated.scrollback": 10000,
  "files.exclude": {
    "**/.git": false
  },
  "security.workspace.trust.enabled": false,
  "basedpyright.analysis.typeCheckingMode": "standard",
  "Lingma.DisplayLanguage": "简体中文",
  "Lingma.PreferredLanguage for AI Chat": "简体中文",
  "Lingma.PreferredLanguage forCommitMessage": "简体中文",
  "git.blame.editorDecoration.enabled": true,
  "rust-analyzer.server.path": "rust-analyzer"
}`
