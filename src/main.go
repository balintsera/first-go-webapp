package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)

const (
  // Endpoint API endpoint base url
  Endpoint = "/api/v/1" 
)

func main() {
  router := httprouter.New()
  
  // User Collection resources
  router.GET(Endpoint + "/users", UserIndex)
  router.POST(Endpoint + "/users", UserCreate)

  // User singular
  router.GET(Endpoint + "/users/:id", UserShow)
  router.PUT(Endpoint + "/users/:id", UserUpdate)

  // Frontpage
  router.GET("/", HomeIndex)

  // Start the server
  println("Starting server on port 8080")
  http.ListenAndServe(":8080", router)
  
}
