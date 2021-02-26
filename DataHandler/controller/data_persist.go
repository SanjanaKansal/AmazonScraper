package controller

import (
	"encoding/json"
	"github.com/SanjanaKansal/data_handler/models"
	"github.com/SanjanaKansal/data_handler/service"
	"github.com/SanjanaKansal/data_handler/utils"
	"log"
	"net/http"
)

// HandleDataPersistence handles the data persisting requests. It expects product data (
//models.ScrapedData) a parameter in the request's body.
func HandleDataPersistence(writer http.ResponseWriter, request *http.Request) {
	log.Println("............HandleDataPersistence Called.............")
	var input models.ScrapedData
	ret := utils.GetStructFromInput(writer, request, &input)
	if ret == false {
		utils.SendEmptyResponse(writer)
		return
	}
	out := service.PersistData(input)
	if !out.Success {
		log.Printf("Failed to write data in file with error: %s\n", out.Message)
		json.NewEncoder(writer).Encode(out)
		return
	}
	json.NewEncoder(writer).Encode(out)
}
