package commandline

import (
	"fmt"
	"github.com/evgeniylapta/go/app/village"
	"github.com/evgeniylapta/go/app/worldmap"
)

const GenWorldProgParam = "generate-world"

func WorldGenStart() {
	vils := village.Generate(100)
	worldMap := worldmap.New(1000, vils)

	fmt.Println(worldMap)
}
