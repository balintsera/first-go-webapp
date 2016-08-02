package main

import (
  "encoding/json"
  "net/http"
)

// JSONError error object to send as a response
type JSONError struct {
  Status string
  HTTPErrorCode int
  Message string 
} 

// Send creates a json object from JSONError object
// and sends to the client
func (object *JSONError) Send(response http.ResponseWriter) (err error) {
  response.Header().Set("Content-Type", "application/json")
  jsonResponse, err := json.Marshal(object)
  response.WriteHeader(object.HTTPErrorCode)
  response.Write(jsonResponse)
  return err
}