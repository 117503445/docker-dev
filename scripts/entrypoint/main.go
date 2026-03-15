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
	log.Info().Msg("Starting entrypoint")

	// replaceCodeServerAppName := func() {
	// 	// replace /usr/lib/code-server/out/node/routes/vscode.js
	// 	// name: appName -> name: "vsc"
	// 	f := "/usr/lib/code-server/out/node/routes/vscode.js"

	// 	old := "name: appName"
	// 	new := `name: "vsc"`

	// 	log.Info().
	// 		Str("file", f).
	// 		Str("old", old).
	// 		Str("new", new).
	// 		Msg("Replace code-server app name")
	// 	if !goutils.FileExists(f) {
	// 		log.Warn().Str("file", f).Msg("vscode.js not exists")
	// 		return
	// 	}
	// 	content, err := goutils.ReadText(f)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Failed to read code-server config file")
	// 		return
	// 	}
	// 	content = strings.ReplaceAll(content, "name: appName", `name: "vsc"`)
	// 	err = goutils.WriteText(f, content)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Failed to write code-server config file")
	// 		return
	// 	}
	// }
	// replaceCodeServerAppName()

	// 检测 code-server 配置路径 (兼容 linuxserver 镜像的 /config 目录)
	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	if goutils.PathExists("/config") {
		codeServerConfigPath = "/config/.config/code-server/config.yaml"
	}
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
	// 确保目录存在
	configDir := "/root/.config/code-server"
	if goutils.PathExists("/config") {
		configDir = "/config/.config/code-server"
	}
	if !goutils.PathExists(configDir) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			log.Panic().Err(err).Msg("Failed to create code-server config directory")
		}
	}
	// 使用 os.WriteFile 直接写入，确保文件权限正确
	err := os.WriteFile(codeServerConfigPath, []byte(codeServerConfigText), 0644)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to write code-server config file")
	}
	log.Info().Str("path", codeServerConfigPath).Msg("Wrote code-server config file")

	// 检查 code-server 是否已经在运行
	codeServerAlreadyRunning := false
	if goutils.FileExists("/proc") {
		// 检查是否有 code-server 进程在监听指定端口
		cmd := exec.Command("ss", "-tlnp")
		output, err := cmd.Output()
		if err == nil && strings.Contains(string(output), ":"+codeServerPort) {
			log.Info().Str("port", codeServerPort).Msg("Port already in use, checking if code-server is running")
			// 检查是否是 code-server 在使用这个端口
			if strings.Contains(string(output), "code-server") || strings.Contains(string(output), "node") {
				log.Info().Msg("code-server appears to be already running, skipping start")
				codeServerAlreadyRunning = true
			}
		}
	}

	if !codeServerAlreadyRunning {
		go func() {
			fileLog := "/docker-dev/logs/code-server.log"
			file, err := os.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Error().Err(err).Msg("Failed to open log file, code-server will not start")
				return
			}
			defer file.Close()

			// 使用 --config 参数指定配置文件路径
			cmd := exec.Command("/usr/sbin/code-server", "--config", codeServerConfigPath)
			// cmd.Stdin = os.Stdin
			cmd.Stdout = file
			cmd.Stderr = file
			cmd.Dir = "/docker-dev"
			cmd.Env = os.Environ()
			log.Info().Str("log", fileLog).Str("config", codeServerConfigPath).Msg("Starting code-server")
			if err := cmd.Run(); err != nil {
				log.Error().Err(err).Msg("Failed to run code-server")
			}
		}()
	}

	go func() {
		// sshd -D 使其在前台运行，不会 fork 后退出
		args := []string{"-D"}
		if os.Getenv("SSHD_PORT") != "" {
			args = append(args, "-p", os.Getenv("SSHD_PORT"))
		}
		cmd := exec.Command("/usr/sbin/sshd", args...)
		log.Info().
			Str("cmd", cmd.String()).
			Msg("Starting sshd in foreground mode")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Failed to run sshd")
		}
	}()

	// 启动 sshole agent
	go func() {
		hubServer := os.Getenv("SSHOLE_AGENT_HUB_SERVER")
		if hubServer == "" {
			log.Info().Msg("SSHOLE_AGENT_HUB_SERVER is not set, skipping sshole agent")
			return
		}
		logFile := "/docker-dev/logs/sshole-agent.log"
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Error().Err(err).Msg("Failed to open sshole agent log file")
			return
		}
		defer file.Close()

		args := []string{"--hub-server", hubServer}
		if auth := os.Getenv("SSHOLE_AGENT_AUTH"); auth != "" {
			args = append(args, "--auth", auth)
		}
		if name := os.Getenv("SSHOLE_AGENT_NAME"); name != "" {
			args = append(args, "--name", name)
		}
		if localPort := os.Getenv("SSHOLE_AGENT_LOCAL_PORT"); localPort != "" {
			args = append(args, "--local-port", localPort)
		}
		if skipSshd := os.Getenv("SSHOLE_AGENT_SKIP_SSHD"); skipSshd != "" {
			args = append(args, "--skip-sshd")
		}

		cmd := exec.Command("agent", args...)
		cmd.Stdout = file
		cmd.Stderr = file
		cmd.Env = os.Environ()
		log.Info().Strs("args", args).Str("log", logFile).Msg("Starting sshole agent")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("Failed to run sshole agent")
		}
	}()

	// go func() {
	// 	// if /init exists, run it
	// 	if goutils.PathExists("/init") {
	// 		log.Info().Msg("Running /init")
	// 		result, err := goutils.Exec("/init")
	// 		if err != nil {
	// 			log.Error().Err(err).Msg("Failed to run /init")
	// 		}else{
	// 			log.Info().Interface("result", result).Msg("/init finished")
	// 		}
	// 	}
	// }()

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
		log.Panic().Msg("Cygwin terminal is not supported")
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
		// 保持进程运行，防止 s6-overlay 不断重启
		log.Info().Msg("Running in non-TTY mode, waiting forever...")
		select {}
	}
}
