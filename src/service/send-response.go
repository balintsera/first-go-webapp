package service

import (
	"encoding/json"
	"net/http"
)

type jsonResponseSourceObject interface{}

func SendJSONResponse(object jsonResponseSourceObject, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(object)
	if err != nil {
		panic("JSON conversion failed")
	}
	response.Write(jsonResponse)
}
