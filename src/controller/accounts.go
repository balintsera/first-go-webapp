package controller

import (
	"net/http"

	"twitter-epub/src/model"
	"twitter-epub/src/service"

	"github.com/julienschmidt/httprouter"
)

// AccountsIndex lists specified user's accounts
func AccountsIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("AccountsIndex")
	var user model.User
	user, err := user.Find(routeParams.ByName("id"))
	if err != nil {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusNotFound,
			Message:       "No user found",
		}
		JSONResponse.Send(response)
		return
	}

	// Send user list
	service.SendJSONResponse(user.Accounts, response)
}

// AccountsCreate creates an account to the specified user it belongs
func AccountsCreate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("AccountsCreate")
}
