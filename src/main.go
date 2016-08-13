package main

import (
	"net/http"
	"twitter-epub/src/controller"

	"github.com/julienschmidt/httprouter"
)

const (
	// Endpoint API endpoint base url
	Endpoint = "/api/v/1"
)

func main() {
	router := httprouter.New()

	// Users Collection resources
	router.GET(Endpoint+"/users", controller.UsersIndex)
	router.POST(Endpoint+"/users", controller.UsersCreate)

	// Users singular
	router.GET(Endpoint+"/users/:id", controller.UsersShow)
	router.PUT(Endpoint+"/users/:id", controller.UsersUpdate)
	router.DELETE(Endpoint+"/users/:id", controller.UsersDelete)

	// Users.accounts
	router.GET(Endpoint+"/users/:id/accounts", controller.AccountsIndex)
	router.POST(Endpoint+"/users/:id/accounts", controller.AccountsCreate)

	// Frontpage
	router.GET("/", HomeIndex)

	// Start the server
	println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)

}
