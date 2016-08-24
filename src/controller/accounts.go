package controller

import (
	"fmt"
	"net/http"
	"strings"

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

	request.ParseForm()

	// Validation: mandatory fields
	mandatoryFields := []string{"type", "title", "mail"}
	if missings, ok := missingMandatory(mandatoryFields, request); !ok {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       fmt.Sprintf("%s parameter is mandatory", strings.Join(missings, ",")),
		}
		JSONResponse.Send(response)
		return
	}

	//ID, Title, service
	validator := service.Validation{FieldType: service.FieldTypeMail, Value: request.Form.Get("mail")}
	validator.Run()
	if !validator.Valid {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       "'mail' parameter is not valid. Please use a valid mail address. You sent: " + request.Form.Get("mail"),
		}
		JSONResponse.Send(response)
		return
	}

}

// checkMandatory checks existence of the fields of the first param in the form in the second
func missingMandatory(fields []string, request *http.Request) (missingFields []string, ok bool) {
	for _, field := range fields {
		// checking input length
		if len(request.Form.Get(field)) < 1 {
			// Not set, send error
			missingFields = append(missingFields, field)
		}
	}
	// If there's no field in the array, it means there's no missing fields
	if len(missingFields) < 1 {
		return missingFields, true
	}

	return missingFields, false
}
