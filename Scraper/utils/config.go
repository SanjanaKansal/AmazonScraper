package utils

import (
	"encoding/json"
	"github.com/SanjanaKansal/scraper/models"
	"io/ioutil"
	"log"
	"os"
)

var configData models.ConfigJson

// Initialize the configdata from the configuration file.
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