package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  "fmt"
)

type jsonResponseSourceObject interface {}

// UserIndex lists users
func UserIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserIndex")
}

// UserCreate create a new user
func UserCreate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserCreate")
}

// UserShow display a user
func UserShow(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserShow")
  var user User
  user, err := user.Find(routeParams.ByName("id"))
  if err != nil {
    // send http error header
	  JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusNotFound, 
      Message: "User not found",
    }
    err := JSONResponse.Send(response)
    if err != nil {
      // unknown error
    }
    return
  }

  // Send response
  if err != nil {
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusInternalServerError, 
      Message: "Unknown error occured when trying to convert user object to json",
    }
    JSONResponse.Send(response)
    return
  }
  sendJSONResponse(user, response)
}

// UserUpdate update a user
func UserUpdate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserUpdate")
}

func sendJSONResponse(object jsonResponseSourceObject, response http.ResponseWriter) {
  fmt.Printf("%+v", object)
  response.Header().Set("Content-Type", "application/json")
  jsonResponse, err := json.Marshal(object)
  if err != nil {
    panic("JSON conversion failed")
  }
  response.Write(jsonResponse)
}
