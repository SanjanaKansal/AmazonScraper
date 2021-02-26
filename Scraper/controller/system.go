package controller

import (
	"log"
	"net/http"
)
// Handler for healthcheck requests.
func HandlerHealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	log.Println("............Healthcheck success.............")
	return
}
