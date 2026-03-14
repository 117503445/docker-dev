package main

import (
	"embed"
	"os"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

//go:embed settings.json
var settingsFS embed.FS

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Starting vibe-init")

	home := os.Getenv("HOME")
	if home == "" {
		home = "/root"
	}

	// Create .claude directory
	claudeDir := home + "/.claude"
	if !goutils.PathExists(claudeDir) {
		log.Info().Str("path", claudeDir).Msg("Creating .claude directory")
		if err := os.MkdirAll(claudeDir, 0755); err != nil {
			log.Fatal().Err(err).Msg("Failed to create .claude directory")
		}
	}

	// Write embedded Claude settings
	settingsPath := claudeDir + "/settings.json"
	log.Info().Str("path", settingsPath).Msg("Writing embedded Claude settings")
	content, err := settingsFS.ReadFile("settings.json")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read embedded settings.json")
	}
	err = goutils.WriteText(settingsPath, string(content))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write Claude settings")
	}

	log.Info().Msg("vibe-init completed")
}