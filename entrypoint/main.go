package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "embed"
	"github.com/117503445/goutils"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog/log"
)

func installVscExtensions() {
	// list all dirs in /root/.vscode-server/bin
	vscVersions, err := filepath.Glob("/root/.vscode-server/bin/*")
	if err != nil {
		log.Error().Err(err).Msg("Failed to list Visual Studio Code Server versions")
		return
	}
	if len(vscVersions) == 0 {
		log.Warn().Msg("No Visual Studio Code Server version found, skip installing extensions")
		return
	}

	fileVsc := fmt.Sprintf("/root/.vscode-server/bin/%s/bin/code-server", filepath.Base(vscVersions[0]))

	vscExts := os.Getenv("VSC_EXTS")
	exts := strings.Split(vscExts, ",")
	for _, ext := range exts {
		cmd := fmt.Sprintf("%s --install-extension %s", fileVsc, ext)
		// log.Debug().Str("ext", ext).Msg("Install Visual Studio Code Extension")
		goutils.Exec(cmd)
		// log.Debug().Str("ext", ext).Msg("Install Visual Studio Code Extension Done")
	}

	cmd := fmt.Sprintf("%s --update-extensions", fileVsc)

	goutils.Exec(cmd)
}

//go:embed code-server-config-template.yaml
var codeServerConfigTemplate string

func main() {
	goutils.InitZeroLog()

	installVscExtensions()

	// log.Debug().Msg("Update Arch Linux Packages")
	// goutils.CMD("", "pacman", "-Syu", "--noconfirm")
	goutils.Exec("pacman -Syu --noconfirm")
	// log.Debug().Msg("Update Arch Linux Packages Done")

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
		// goutils.CMD("", "tail", "-f", "/dev/null")
		goutils.Exec("tail -f /dev/null")
	}
}
