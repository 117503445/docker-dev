package main

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/117503445/goutils"
	"github.com/rs/zerolog/log"
)

//go:embed AGENTS.md claude/settings.json codex/hooks.json hooks/vibe_hook.js all:skills
var embedFS embed.FS

func main() {
	goutils.InitZeroLog()
	log.Info().Msg("Starting vibe-init")

	home := os.Getenv("HOME")
	if home == "" {
		home = "/root"
	}

	aiDir := filepath.Join(home, ".ai")
	mustMkdirAll(aiDir)
	mustCopyFile("hooks/vibe_hook.js", filepath.Join(aiDir, "vibe_hook.js"), 0755)

	claudeDir := home + "/.claude"
	mustMkdirAll(claudeDir)
	mustCopyFile("claude/settings.json", filepath.Join(claudeDir, "settings.json"), 0644)
	mustCopyFile("AGENTS.md", filepath.Join(claudeDir, "AGENTS.md"), 0644)
	mustWriteText(filepath.Join(claudeDir, "CLAUDE.md"), "@~/.claude/AGENTS.md\n", 0644)
	mustCopyFS("skills", filepath.Join(claudeDir, "skills"))

	codexDir := home + "/.codex"
	mustMkdirAll(codexDir)
	mustCopyFile("AGENTS.md", filepath.Join(codexDir, "AGENTS.md"), 0644)
	mustCopyFile("codex/hooks.json", filepath.Join(codexDir, "hooks.json"), 0644)
	mustCopyFS("skills", filepath.Join(codexDir, "skills"))
	ensureCodexHooksEnabled(filepath.Join(codexDir, "config.toml"))

	log.Info().Msg("vibe-init completed")
}

func mustMkdirAll(path string) {
	if !goutils.PathExists(path) {
		log.Info().Str("path", path).Msg("Creating directory")
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Error().Err(err).Msg("Failed to create directory")
		}
	}
}

func mustCopyFile(srcPath string, destPath string, perm fs.FileMode) {
	log.Info().Str("src", srcPath).Str("dest", destPath).Msg("Writing embedded file")
	content, err := embedFS.ReadFile(srcPath)
	if err != nil {
		log.Error().Err(err).Str("path", srcPath).Msg("Failed to read embedded file")
		return
	}
	mustWriteText(destPath, string(content), perm)
}

func mustWriteText(path string, content string, perm fs.FileMode) {
	if err := os.WriteFile(path, []byte(content), perm); err != nil {
		log.Error().Err(err).Str("path", path).Msg("Failed to write file")
	}
}

func mustCopyFS(srcDir, destDir string) {
	log.Info().Str("src", srcDir).Str("dest", destDir).Msg("Writing embedded directory")
	if err := copyFS(srcDir, destDir); err != nil {
		log.Error().Err(err).Str("path", destDir).Msg("Failed to write directory")
	}
}

func copyFS(srcDir, destDir string) error {
	return fs.WalkDir(embedFS, srcDir, func(path string, d fs.DirEntry, err error) error {
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

		data, err := embedFS.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(destPath, data, 0644)
	})
}

func ensureCodexHooksEnabled(configPath string) {
	data, err := os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		log.Error().Err(err).Str("path", configPath).Msg("Failed to read Codex config")
		return
	}

	content := string(data)
	if strings.Contains(content, "codex_hooks") {
		return
	}

	if strings.Contains(content, "[features]") {
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if strings.TrimSpace(line) == "[features]" {
				lines = append(lines[:i+1], append([]string{"codex_hooks = true"}, lines[i+1:]...)...)
				mustWriteText(configPath, strings.Join(lines, "\n"), 0644)
				return
			}
		}
	}

	if strings.TrimSpace(content) != "" && !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	content += "\n[features]\ncodex_hooks = true\n"
	mustWriteText(configPath, content, 0644)
}
