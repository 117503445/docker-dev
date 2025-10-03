package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "embed"

	"github.com/117503445/goutils"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

const codeServerConfigTemplate string = `bind-addr: 0.0.0.0:%s
auth: password
password: %s
cert: false`

func main() {
	goutils.InitZeroLog()
	goutils.ExecOpt.DumpOutput = true

	replaceCodeServerAppName := func() {
		// replace /usr/lib/code-server/out/node/routes/vscode.js
		// name: appName -> name: "vsc"
		f := "/usr/lib/code-server/out/node/routes/vscode.js"

		old := "name: appName"
		new := `name: "vsc"`

		log.Info().
			Str("file", f).
			Str("old", old).
			Str("new", new).
			Msg("Replace code-server app name")
		if !goutils.FileExists(f) {
			log.Warn().Str("file", f).Msg("vscode.js not exists")
			return
		}
		content, err := goutils.ReadText(f)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read code-server config file")
			return
		}
		content = strings.ReplaceAll(content, "name: appName", `name: "vsc"`)
		err = goutils.WriteText(f, content)
		if err != nil {
			log.Error().Err(err).Msg("Failed to write code-server config file")
			return
		}
	}
	replaceCodeServerAppName()

	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	codeServerPassword := os.Getenv("CODE_SERVER_PASSWORD")
	if codeServerPassword == "" {
		log.Warn().Msg("CODE_SERVER_PASSWORD is not set, use default password")
		codeServerPassword = "123456"
	}
	codeServerPort := os.Getenv("CODE_SERVER_PORT")
	if codeServerPort == "" {
		log.Warn().Msg("CODE_SERVER_PORT is not set, use default port 4444")
		codeServerPort = "4444"
	}
	codeServerConfigText := fmt.Sprintf(codeServerConfigTemplate, codeServerPort, codeServerPassword)
	err := goutils.WriteText(codeServerConfigPath, codeServerConfigText)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write code-server config file")
	}

	go func() {
		fileLog := "/docker-dev/logs/code-server.log"
		file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open file")
			return
		}
		defer file.Close()

		cmd := exec.Command("/usr/sbin/code-server")
		// cmd.Stdin = os.Stdin
		cmd.Stdout = file
		cmd.Stderr = file
		cmd.Dir = "/docker-dev"
		cmd.Env = os.Environ()
		log.Info().Str("log", fileLog).Msg("Starting code-server")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Failed to run code-server")
		}
	}()

	go func() {
		cmd := exec.Command("/usr/sbin/sshd")
		log.Info().Msg("Starting sshd")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Failed to run sshd")
		}
	}()

	fileCustomEntrypoint := "/entrypoint"
	if goutils.PathExists(fileCustomEntrypoint) {
		env := os.Environ()
		envOption := goutils.WithEnv{}
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			env = append(env, pair[0]+"="+pair[1])
		}

		goutils.Exec(fileCustomEntrypoint, envOption)
	}

	var isTTY bool
	if isatty.IsTerminal(os.Stdout.Fd()) {
		isTTY = true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		log.Fatal().Msg("Cygwin terminal is not supported")
	} else {
		isTTY = false
	}

	// log.Debug().Bool("isTTY", isTTY).Msg("Check if TTY")

	if isTTY {
		cmd := exec.Command("/bin/zsh")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		cmd.Env = os.Environ()
		// log.Debug().Msg("Enter zsh shell")
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("Failed to run zsh shell")
		}
		// log.Debug().Msg("Exit zsh shell")
	} else {
		goutils.Exec("tail -f /dev/null")
	}
}
