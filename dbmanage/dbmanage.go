package dbmanage

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
)

type dbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func getAppConf() dbConfig {
	jsonFile, openErr := os.Open("dbconfig.json")

	if openErr != nil {
		fmt.Println(openErr)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var appConf dbConfig
	unmarshalErr := json.Unmarshal(byteValue, &appConf)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	return appConf
}

func ConnectToDb() *gorm.DB {
	appConf := getAppConf()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConf.Host, appConf.Port, appConf.User, appConf.Password, appConf.Dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
