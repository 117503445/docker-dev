package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/mattn/go-isatty"
	"os/exec"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	log.Debug().Msg("Update Arch Linux Packages")
	CMD("", "pacman", "-Syu", "--noconfirm")
	log.Debug().Msg("Update Arch Linux Packages Done")

	fileCustomEntrypoint := "/entrypoint"
	if _, err := os.Stat(fileCustomEntrypoint); err == nil {
		log.Debug().Msg("Run Custom Entrypoint")
		CMD("", fileCustomEntrypoint)
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
		CMD("", "tail", "-f", "/dev/null")
	}
}

func CMD(cwd string, command string, args ...string) {
	var err error
	if cwd == "" {
		cwd, err = os.Getwd()
		if err != nil {
			log.Error().Err(err).Msg("Failed to get current working directory")
			return
		}
	}
	commandStr := command
	for _, arg := range args {
		commandStr += " " + arg
	}
	log.Debug().Str("cwd", cwd).Str("command", commandStr).Msg("Run Command")
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Error().Err(err).Str("cwd", cwd).Str("command", command).Strs("args", args).
			Msg("Failed to run command")
	}
}
