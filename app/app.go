package app

import (
	"fmt"
	"github.com/evgeniylapta/go/utils"
	"github.com/evgeniylapta/go/village/generator"
	"github.com/evgeniylapta/go/worldmap"
	"time"
)

const generateParameter = "generate"

func generateStart() {
	vils := generator.Generate(100)
	worldMap := worldmap.New(1000, vils)

	fmt.Println(worldMap)
}

func runStart() {
	for {
		time.Sleep(time.Second)
	}
}

func Start() {
	//router := gin.Accounts{}.Default()

	switch  {
	case utils.StringFound(utils.ArgsWithoutProg(), generateParameter):
		generateStart()
	default:
		runStart()
	}
}
