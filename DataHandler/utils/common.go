package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetStructFromInput reads data from the body of a http and sets writer's content-type as application/json.
// Request type object and if there is some data in the body then it parses the JSON-encoded data and stores the result
// in the value pointed to by arg.
func GetStructFromInput(writer http.ResponseWriter, request *http.Request, arg interface{}) bool {
	writer.Header().Set("content-type", "application/json")
	body, err := ioutil.ReadAll(request.Body)

	ret := true
	if len(body) == 0 || err != nil {
		log.Println("Empty body")
		ret = false
		return ret
	} else {
		err = json.Unmarshal(body, arg)
		if err != nil {
			log.Println("Unmarshalling error: ", err)
			ret = false
			return ret
		}
	}
	return ret
}

// SendEmptyResponse sends Bad Request Response.
func SendEmptyResponse(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/text")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.WriteHeader(http.StatusBadRequest)
	log.Println("Output: ", "Badrequest")
}
