package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	_ "embed"

	"github.com/117503445/goutils"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

//go:embed code-server-config-template.yaml
var codeServerConfigTemplate string

func main() {
	goutils.InitZeroLog()
	goutils.ExecOpt.DumpOutput = true

	// goutils.Exec("pacman -Syu --noconfirm")

	codeServerConfigPath := "/root/.config/code-server/config.yaml"
	if _, err := os.Stat(codeServerConfigPath); os.IsNotExist(err) {
		codeServerPassword := os.Getenv("CODE_SERVER_PASSWORD")
		if codeServerPassword == "" {
			log.Warn().Msg("CODE_SERVER_PASSWORD is not set, use default password")
			codeServerPassword = "123456"
		}

		codeServerConfigText := fmt.Sprintf(codeServerConfigTemplate, codeServerPassword)

		if err := os.MkdirAll(filepath.Dir(codeServerConfigPath), 0755); err != nil {
			log.Error().Err(err).Msg("Failed to create code-server config directory")
		} else {
			if err := os.WriteFile(codeServerConfigPath, []byte(codeServerConfigText), 0644); err != nil {
				log.Error().Err(err).Msg("Failed to write code-server config file")
			}
		}
	}
	// goutils.CMD("", "systemctl", "start", "code-server@root")
	goutils.Exec("systemctl start code-server@root")

	fileCustomEntrypoint := "/entrypoint"
	if goutils.PathExists(fileCustomEntrypoint) {
		goutils.Exec(fileCustomEntrypoint)
	}

	var isTTY bool
	if isatty.IsTerminal(os.Stdout.Fd()) {
		isTTY = true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		log.Fatal().Msg("Cygwin terminal is not supported")
	} else {
		isTTY = false
	}

	log.Debug().Bool("isTTY", isTTY).Msg("Check if TTY")

	if isTTY {
		cmd := exec.Command("/bin/fish")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Debug().Msg("Enter fish shell")
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err).Msg("Failed to run fish shell")
		}
		log.Debug().Msg("Exit fish shell")
	} else {
		goutils.Exec("tail -f /dev/null")
	}
}
