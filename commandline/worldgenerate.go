package commandline

import (
	"encoding/json"
	"fmt"
	"github.com/evgeniylapta/go/app/village"
	"github.com/evgeniylapta/go/app/worldmap"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"os"
)

const GenWorldProgramParam = "generate-world"

type appConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func getAppConf() appConfig {
	jsonFile, openErr := os.Open("appconfig.json")

	if openErr != nil {
		fmt.Println(openErr)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var appConf appConfig
	unmarshalErr := json.Unmarshal(byteValue, &appConf)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	return appConf
}

func WorldGenStart() {
	appConf := getAppConf()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConf.Host, appConf.Port, appConf.User, appConf.Password, appConf.Dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

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
