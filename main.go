package main

import (
	"github.com/evgeniylapta/go/app"
	"github.com/evgeniylapta/go/commandline"
	"github.com/evgeniylapta/go/utils"
)

func main() {
	switch {
	case utils.StringFound(utils.ArgsWithoutProg(), commandline.GenWorldProgParam):
		commandline.WorldGenStart()
	default:
		app.Start()
	}
}
