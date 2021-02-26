package controller

import (
	"log"
	"net/http"
)

func HandlerHealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	log.Println("............Healthcheck success.............")
	return
}
