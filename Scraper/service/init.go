package service

import (
	"github.com/SanjanaKansal/scraper/models"
	"github.com/SanjanaKansal/scraper/utils"
	"log"
	"net/http"
	"time"
)

// Initialize variables.
func InitializeVars() {
	models.PORT = utils.GetConfig().PORT
	models.HOST = utils.GetConfig().HOST
	models.DATAPERSISTINGAPIURL = utils.GetConfig().DATAPERSISTINGAPIURL
}

// Utility method to check the status of the machine on which this application is hosted.
func CheckOnlineStatus() {
	for {
		_, err := http.Get("http://google.com/")
		if err != nil {
			log.Println("Host lost connection from internet.")
		}
		time.Sleep(5 * time.Second)
	}
}