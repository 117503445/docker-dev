package main

import (
	"os"
	"os/exec"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Starting shell-init")

	home := os.Getenv("HOME")
	if home == "" {
		home = "/root"
	}

	// Change default shell to zsh
	log.Info().Msg("Changing default shell to zsh")
	cmd := exec.Command("chsh", "-s", "/usr/bin/zsh")
	if err := cmd.Run(); err != nil {
		log.Warn().Err(err).Msg("Failed to change default shell, continuing...")
	}

	// Install oh-my-zsh
	ohMyZshPath := home + "/.oh-my-zsh"
	if !goutils.PathExists(ohMyZshPath) {
		log.Info().Msg("Installing oh-my-zsh")
		installScript := "https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh"
		if os.Getenv("CHINA_MIRROR") != "" {
			installScript = "https://install.ohmyz.sh"
		}
		cmd = exec.Command("sh", "-c", "curl -fsSL "+installScript+" | sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Warn().Err(err).Msg("Failed to install oh-my-zsh, continuing...")
		}
	} else {
		log.Info().Msg("oh-my-zsh already installed")
	}

	// Copy .zshrc from /tmp/.zshrc to home directory
	tmpZshrc := "/tmp/.zshrc"
	homeZshrc := home + "/.zshrc"
	if goutils.PathExists(tmpZshrc) {
		log.Info().Str("src", tmpZshrc).Str("dst", homeZshrc).Msg("Copying .zshrc")
		content, err := goutils.ReadText(tmpZshrc)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read .zshrc from /tmp")
		}
		err = goutils.WriteText(homeZshrc, content)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write .zshrc")
		}
	} else {
		log.Warn().Str("path", tmpZshrc).Msg(".zshrc not found in /tmp, skipping")
	}

	// Create .zshrc-custom if not exists
	zshrcCustomPath := home + "/.zshrc-custom"
	if !goutils.PathExists(zshrcCustomPath) {
		log.Info().Str("path", zshrcCustomPath).Msg("Creating empty .zshrc-custom")
		err := goutils.WriteText(zshrcCustomPath, "")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create .zshrc-custom")
		}
	}

	// Configure histmon
	log.Info().Msg("Configuring histmon")
	cmd = exec.Command("histmon", "install")
	cmd.Env = append(os.Environ(), "TOKEN=T", "ENDPOINT=E")
	if err := cmd.Run(); err != nil {
		log.Warn().Err(err).Msg("Failed to configure histmon, continuing...")
	}

	log.Info().Msg("shell-init completed")
}