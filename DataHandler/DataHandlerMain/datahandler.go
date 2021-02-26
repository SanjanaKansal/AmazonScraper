package main

import (
	"github.com/SanjanaKansal/data_handler/app"
	"github.com/SanjanaKansal/data_handler/models"
	"github.com/SanjanaKansal/data_handler/service"
	"github.com/SanjanaKansal/data_handler/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	log.Println("*****Starting Web Scraper Server*******")

	// Reading configuration from config file
	configPath := "config.json"
	utils.InitConfig(configPath)

	// Initializing global variables
	service.InitializeVars()

	// Service to check online status of Host server
	go service.CheckOnlineStatus()
	time.Sleep(1 * time.Second)

	//Assign host and port values from global variables
	host := models.HOST
	port := models.PORT
	log.Println("Host - ", host)
	log.Println("Port - ", port)

	ServerAddress := host + ":" + strconv.Itoa(port)
	Router := app.Router
	log.Println("Listening @ ", ServerAddress)
	// Starting the server
	err := http.ListenAndServe(ServerAddress, Router)
	if err != nil {
		log.Fatal("failed bringing up the server")
	}
}
