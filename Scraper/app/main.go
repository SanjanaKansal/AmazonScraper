package app

import (
	"github.com/SanjanaKansal/scraper/controller"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/webscraper/healthcheck", controller.HandlerHealthCheck)
	Router.HandleFunc("/webscraper/scrape_page", controller.HandlePageScraping).Methods("POST")
}
