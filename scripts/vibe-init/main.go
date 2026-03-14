package main

import (
	"os"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

// ClaudeSettingsContent is the Claude Code settings
const ClaudeSettingsContent = `{
    "skipDangerousModePermissionPrompt": true,
    "enabledPlugins": {
        "gopls-lsp@claude-plugins-official": true
    },
    "permissions": {
        "allow": [
            "Read(**)",
            "Edit(**)",
            "Bash(*)",
            "LS(**)",
            "Grep(**)",
            "Glob(**)",
            "WebFetch",
            "MCP"
        ],
        "deny": []
    }
}`

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

	// Copy settings.json from /tmp/claude-settings.json if exists, otherwise use embedded
	tmpSettings := "/tmp/claude-settings.json"
	settingsPath := claudeDir + "/settings.json"

	if goutils.PathExists(tmpSettings) {
		log.Info().Str("src", tmpSettings).Str("dst", settingsPath).Msg("Copying Claude settings from /tmp")
		content, err := goutils.ReadText(tmpSettings)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read Claude settings from /tmp")
		}
		err = goutils.WriteText(settingsPath, content)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write Claude settings")
		}
	} else {
		log.Info().Str("path", settingsPath).Msg("Writing embedded Claude settings")
		err := goutils.WriteText(settingsPath, ClaudeSettingsContent)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write Claude settings")
		}
	}

	log.Info().Msg("vibe-init completed")
}