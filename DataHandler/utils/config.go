package utils

import (
	"encoding/json"
	"github.com/SanjanaKansal/data_handler/models"
	"io/ioutil"
	"log"
	"os"
)

var configData models.ConfigJson

func InitConfig(configFile string) {
	log.Println("Config File: ", configFile)
	configFP, err := os.Open(configFile)
	defer configFP.Close()
	if err != nil {
		log.Fatal("Unable to read config file")
	}
	byteVal, _ := ioutil.ReadAll(configFP)
	json.Unmarshal(byteVal, &configData)
}

func GetConfig() models.ConfigJson {
	return configData
}