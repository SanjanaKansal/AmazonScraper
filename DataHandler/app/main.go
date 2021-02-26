package app

import (
	"github.com/SanjanaKansal/data_handler/controller"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/datahelper/healthcheck", controller.HandlerHealthCheck)
	Router.HandleFunc("/datahelper/persist_data", controller.HandleDataPersistence).Methods("POST")
}
