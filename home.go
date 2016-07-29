package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

// HomeRoute controller 
func HomeIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("this is the /");
}

