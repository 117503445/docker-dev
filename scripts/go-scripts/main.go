package main

import (
	"github.com/117503445/goutils"
	"github.com/117503445/goutils/glog"
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
)

func init() {
	glog.InitZeroLog()
}

var dirProjectRoot = func() string {
	d, err := goutils.FindGitRepoRoot()
	if err != nil {
		log.Panic().Err(err).Msg("failed to find git repo root")
	}
	return d
}()

func main() {
	ctx := kong.Parse(&cli)
	log.Info().Interface("cli", cli).Send()
	if err := ctx.Run(); err != nil {
		log.Panic().Err(err).Msg("run failed")
	}
}