package commandline

import (
	"fmt"
	"github.com/evgeniylapta/go/app/village"
	"github.com/evgeniylapta/go/app/worldmap"
	"github.com/evgeniylapta/go/dbmanage"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const GenWorldProgramParam = "generate-world"

func WorldGenStart() {
	db := dbmanage.ConnectToDb()

	db.AutoMigrate(&village.Village{})

	const mapSize = 1000
	vils := village.GenerateVillages(1000, -1000, 1000)
	for _, vil := range vils {
		db.Create(&vil)
	}

	worldMap := worldmap.New(mapSize, vils)
	fmt.Println(worldMap)

	defer db.Close()
}
