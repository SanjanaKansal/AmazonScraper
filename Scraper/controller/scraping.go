package controller

import (
	"bytes"
	"encoding/json"
	"github.com/SanjanaKansal/scraper/models"
	"github.com/SanjanaKansal/scraper/service"
	"github.com/SanjanaKansal/scraper/utils"
	"log"
	"net/http"
)

// HandlePageScraping handles the Page Scraping requests. It expects url (string) a parameter in the request's body.
// After scraping the info from the page, it calls another external API to persist that data in files.
func HandlePageScraping(writer http.ResponseWriter, request *http.Request) {
	log.Println("............HandlePageScraping Called.............")
	var input models.ScrapePageInp
	ret := utils.GetStructFromInput(writer, request, &input)
	if ret == false {
		utils.SendEmptyResponse(writer)
		return
	}
	if len(input.URL) == 0{
		utils.SendEmptyResponse(writer)
		return
	}
	out, data := service.ScrapePage(input.URL)
	if !out.Success {
		log.Printf("The scraping failed with error %s\n", out.Message)
		json.NewEncoder(writer).Encode(out)
		return
	}

	// Calling external API for persisting the data in file.
	apiPayload, _ := json.Marshal(data)
	dataPersistingAPI := models.DATAPERSISTINGAPIURL
	log.Printf("Making external API call to persist the scraped data at: %s\n", dataPersistingAPI)
	_, err := http.Post(dataPersistingAPI, "application/json", bytes.NewBuffer(apiPayload))
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		out.Success = false
		out.Message = "HTTP request to persist data failed"
		json.NewEncoder(writer).Encode(out)
		return
	}
	out.Message = "Successfully scraped and persisted data in mongo."
	json.NewEncoder(writer).Encode(out)
}
