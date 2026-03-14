package main

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

//go:embed settings.json CLAUDE.md all:skills
var embedFS embed.FS

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
			log.Error().Err(err).Msg("Failed to create .claude directory")
		}
	}

	// Write embedded Claude settings
	settingsPath := claudeDir + "/settings.json"
	log.Info().Str("path", settingsPath).Msg("Writing embedded Claude settings")
	content, err := embedFS.ReadFile("settings.json")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read embedded settings.json")
	}
	err = goutils.WriteText(settingsPath, string(content))
	if err != nil {
		log.Error().Err(err).Msg("Failed to write Claude settings")
	}

	// Write embedded CLAUDE.md
	claudeMdPath := claudeDir + "/CLAUDE.md"
	log.Info().Str("path", claudeMdPath).Msg("Writing embedded CLAUDE.md")
	claudeMdContent, err := embedFS.ReadFile("CLAUDE.md")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read embedded CLAUDE.md")
	}
	err = goutils.WriteText(claudeMdPath, string(claudeMdContent))
	if err != nil {
		log.Error().Err(err).Msg("Failed to write CLAUDE.md")
	}

	// Write embedded skills directory
	skillsDir := claudeDir + "/skills"
	log.Info().Str("path", skillsDir).Msg("Writing embedded skills")
	err = copyFS(embedFS, "skills", skillsDir)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write skills directory")
	}

	log.Info().Msg("vibe-init completed")
}

// copyFS recursively copies a directory from embed.FS to the filesystem
func copyFS(efs embed.FS, srcDir, destDir string) error {
	return fs.WalkDir(efs, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		data, err := efs.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(destPath, data, 0644)
	})
}