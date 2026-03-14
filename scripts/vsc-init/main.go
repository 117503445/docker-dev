package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/vsc-init/pkg/assets"
	"github.com/117503445/vsc-init/pkg/ext"
	"github.com/rs/zerolog/log"
)

func main() {
	goutils.InitZeroLog()

	err := goutils.WriteText("/root/.local/share/code-server/User/settings.json", assets.Settings)
	if err != nil {
		log.Fatal().Err(err).Msg("write settings.json error")
	}

	err = goutils.WriteText("/root/.local/share/code-server/User/keybindings.json", assets.KeyBindings)
	if err != nil {
		log.Fatal().Err(err).Msg("write keybindings.json error")
	}

	ext.InstallLatestExts()
}
