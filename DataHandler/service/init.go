package service

import (
	"github.com/SanjanaKansal/data_handler/models"
	"github.com/SanjanaKansal/data_handler/utils"
	"log"
	"net/http"
	"time"
)

func InitializeVars() {
	models.PORT = utils.GetConfig().PORT
	models.HOST = utils.GetConfig().HOST
	models.MONGODBNAME = utils.GetConfig().MONGODBNAME
	models.MONGOCOLLECTIONNAME = utils.GetConfig().MONGOCOLLECTIONNAME
	models.MONGOURI = utils.GetConfig().MONGOURI
}

func CheckOnlineStatus() {
	for {
		_, err := http.Get("http://google.com/")
		if err != nil {
			log.Println("Host lost connection from internet.")
		}
		time.Sleep(5 * time.Second)
	}
}