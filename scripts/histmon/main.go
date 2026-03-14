package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/117503445/goutils"
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
)

var cli struct {
	Command    string `env:"COMMAND"`
	Output     string `env:"OUTPUT"`
	StartAt    int    `env:"START_AT"` // 毫秒时间戳
	EndAt      int    `env:"END_AT"`   // 毫秒时间戳
	ExitStatus int    `env:"EXIT_STATUS"`

	Token    string `env:"TOKEN,HISTMON_TOKEN"`
	Endpoint string `env:"ENDPOINT,HISTMON_ENDPOINT"`

	Install struct {
	} `cmd:"" help:"Install to shell"`
	Send struct {
	} `cmd:"" help:"Send message" default:"1"`
}

func init() {
	// 设置全局时区为 UTC+8
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("UTC+8", 8*60*60)
	}
	time.Local = loc
}

func install() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to get home directory")
	}

	fileZshrc := path.Join(home, ".zshrc")
	if !goutils.FileExists(fileZshrc) {
		log.Panic().
			Str("file", fileZshrc).
			Msg("File .zshrc not found")
	}

	if cli.Endpoint == "" {
		log.Panic().Msg("Please set ENDPOINT environment variable")
	}

	if cli.Token == "" {
		log.Panic().Msg("Please set TOKEN environment variable")
	}

	text, err := goutils.ReadText(fileZshrc)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to read .zshrc")
	}

	// 如果不存在 source ~/.zsh/histmon.zsh 这一行，就添加进 .zshrc
	sourceLine := "source ~/.zsh/histmon.zsh"
	if !strings.Contains(text, sourceLine) {
		text += "\n" + sourceLine + "\n"
		err = goutils.WriteText(fileZshrc, text)
		if err != nil {
			log.Panic().Err(err).Msg("Failed to write .zshrc")
		}
	}

	// 创建 .zsh 目录
	zshDir := path.Join(home, ".zsh")
	err = os.MkdirAll(zshDir, 0755)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create .zsh directory")
	}

	// 使用 text/template 来处理 histmon.zsh 文件模板
	histmonScriptTemplate := `autoload -Uz add-zsh-hook

typeset -g zsh_command_start_time
typeset -g zsh_current_command

preexec() {
    zsh_command_start_time=$(date +%s%3N)  # 毫秒时间戳
    zsh_current_command=$1
}

precmd() {
    local exit_status=$?
    local end_time=$(date +%s%3N)  # 毫秒时间戳
    
    if [[ -n "$zsh_command_start_time" && -n "$zsh_current_command" ]]; then
        # 调用 histmon，重定向所有输出到 /dev/null
        (COMMAND="$zsh_current_command" \
        START_AT="$zsh_command_start_time" \
        END_AT="$end_time" \
        EXIT_STATUS="$exit_status" \
        TOKEN="{{.Token}}" \
        ENDPOINT="{{.Endpoint}}" \
        histmon >/dev/null 2>&1 &)
        
        # 清理变量
        unset zsh_command_start_time
        unset zsh_current_command
    fi
}
`

	tmpl, err := template.New("histmon").Parse(histmonScriptTemplate)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to parse histmon script template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		Token    string
		Endpoint string
	}{
		Token:    cli.Token,
		Endpoint: cli.Endpoint,
	})
	if err != nil {
		log.Panic().Err(err).Msg("Failed to execute histmon script template")
	}

	histmonFile := path.Join(zshDir, "histmon.zsh")
	err = goutils.WriteText(histmonFile, buf.String())
	if err != nil {
		log.Panic().Err(err).Msg("Failed to write histmon.zsh")
	}

	log.Info().Msg("Successfully installed histmon to your shell")
}

func send() {
	// 如果执行时间小于 5 秒，则不发送消息
	if cli.EndAt-cli.StartAt < 5000 {
		log.Info().Msg("Command execution time is less than 5 second, no message will be sent")
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to get hostname")
	}

	content := fmt.Sprintf(`# 命令执行完成
- **主机名**: %s
- **命令**: %s
- **退出状态码**: %d
- **开始时间**: %s
- **结束时间**: %s
- **执行时长**: %s`,
		hostname,
		cli.Command,
		cli.ExitStatus,
		time.UnixMilli(int64(cli.StartAt)).Format("2006-01-02 15:04:05"),
		time.UnixMilli(int64(cli.EndAt)).Format("2006-01-02 15:04:05"),
		goutils.DurationToStr(time.Duration(cli.EndAt-cli.StartAt)*time.Millisecond),
	)

	{
		payload := map[string]any{
			"title":       "命令执行完成",
			"description": "命令执行完成",
			"content":     content,
			"channel":     "ding",
			"token":       cli.Token,
		}

		// 将数据转换为 JSON
		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Panic().Err(err).Msg("Failed to marshal JSON")
			return
		}

		// 创建 HTTP 请求
		resp, err := http.Post(
			cli.Endpoint,
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			log.Panic().Err(err).Msg("Failed to send message")
			return
		}
		defer resp.Body.Close()
	}
}

func main() {
	goutils.InitZeroLog()

	ctx := kong.Parse(&cli)

	log.Info().Interface("cli", cli).Msg("Starting histmon")
	switch ctx.Command() {
	case "install":
		install()
	case "send":
		send()
	default:
		panic(ctx.Command())
	}

}
