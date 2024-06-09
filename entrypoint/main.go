package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os/exec"

	"github.com/mattn/go-isatty"
	"github.com/117503445/goutils"
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
		log.Debug().Str("ext", ext).Msg("Install Visual Studio Code Extension")
		goutils.CMD("", fileVsc, "--install-extension", ext)
		log.Debug().Str("ext", ext).Msg("Install Visual Studio Code Extension Done")
	}

	goutils.CMD("", fileVsc, "--update-extensions")
}

func main() {
	goutils.InitZeroLog()

	installVscExtensions()

	log.Debug().Msg("Update Arch Linux Packages")
	goutils.CMD("", "pacman", "-Syu", "--noconfirm")
	log.Debug().Msg("Update Arch Linux Packages Done")

	fileCustomEntrypoint := "/entrypoint"
	if _, err := os.Stat(fileCustomEntrypoint); err == nil {
		log.Debug().Msg("Run Custom Entrypoint")
		goutils.CMD("", fileCustomEntrypoint)
		log.Debug().Msg("Run Custom Entrypoint Done")
	}

	var isTTY bool
	if isatty.IsTerminal(os.Stdout.Fd()) {
		isTTY = true
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		panic("Cygwin terminal is not supported")
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
		goutils.CMD("", "tail", "-f", "/dev/null")
	}
}
