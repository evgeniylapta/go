package main

import (
	"github.com/evgeniylapta/go/api"
	"github.com/evgeniylapta/go/commandline"
	"github.com/evgeniylapta/go/utils"
)

func shouldStartWorldGenerate() bool {
	return utils.StringFound(utils.ArgsWithoutProg(), commandline.GenWorldProgramParam)
}

func main() {
	switch {
	case shouldStartWorldGenerate():
		commandline.WorldGenStart()
	default:
		api.Start()
	}
}
