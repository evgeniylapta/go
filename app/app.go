package app

import (
	"fmt"
	"github.com/evgeniylapta/go/utils"
	"github.com/evgeniylapta/go/village/generator"
	"github.com/evgeniylapta/go/worldmap"
	"github.com/gin-gonic/gin"
)

const generateParameter = "generate"

func generateStart() {
	vils := generator.Generate(100)
	worldMap := worldmap.New(1000, vils)

	fmt.Println(worldMap)
}

func runStart() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":4300")
}

func Start() {
	switch {
	case utils.StringFound(utils.ArgsWithoutProg(), generateParameter):
		generateStart()
	default:
		runStart()
	}
}
