package controller

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"twitter-epub/src/model"
	"twitter-epub/src/service"

	"github.com/julienschmidt/httprouter"
)

// UsersIndex lists Userss
func UsersIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("UsersIndex")
	var user model.User
	users, err := user.FindAll()
	if err != nil {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusNotFound,
			Message:       "No Users found",
		}
		JSONResponse.Send(response)
		return
	}

	// Send Users list
	service.SendJSONResponse(users, response)
}

// UsersDelete deletes a user via Delete
func UsersDelete(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("UsersDelete: " + routeParams.ByName("id"))
	// validate id?
	user := model.User{}
	user, err := user.Find(routeParams.ByName("id"))
	if err != nil {
		// Users not found or some other error, send error
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       fmt.Sprintf(`Users not found by this id %s`, routeParams.ByName("id")),
		}
		JSONResponse.Send(response)
		return
	}
	err = user.Remove()
	if err != nil {
		// Users not found or some other error, send error
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       fmt.Sprintf(`Error when deleting Users. Id: %s`, routeParams.ByName("id")),
		}
		JSONResponse.Send(response)
		return
	}
	service.SendJSONResponse(user, response)
	return
}

// UsersCreate create a new Users via POST
func UsersCreate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("UsersCreate")

	request.ParseForm()

	// Validation: is mail field set?
	if len(request.Form.Get("mail")) < 1 {
		// Not set, send error
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       "'mail' parameter is mandatory",
		}
		JSONResponse.Send(response)
		return
	}

	// Validate email via regexp
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

	// Check for existence
	found, _ := model.GetUserByMail(request.Form.Get("mail"))
	fmt.Printf("found mail: %+v", found)
	// err is true if not found
	if len(found.Mail) > 0 {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusBadRequest,
			Message:       "A Users is already registered with this mail",
		}
		JSONResponse.Send(response)
		return
	}

	fmt.Printf("mail %+v", request.Form.Get("mail"))

	// Creation
	user := model.User{Mail: request.Form.Get("mail")}
	user.GenerateID()
	user.Insert()

	// @TODO insert to database
	fmt.Printf("new user: %+v", user)
	service.SendJSONResponse(user, response)
}

// UsersShow display a Users
func UsersShow(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	println("UsersShow")
	var user model.User
	user, err := user.Find(routeParams.ByName("id"))
	if err != nil {
		// send http error header
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusNotFound,
			Message:       "Users not found",
		}
		err := JSONResponse.Send(response)
		if err != nil {
			// unknown error
		}
		return
	}

	// Send response
	if err != nil {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusInternalServerError,
			Message:       "Unknown error occured when trying to convert Users object to json",
		}
		JSONResponse.Send(response)
		return
	}
	service.SendJSONResponse(user, response)
}

// UsersUpdate update a Users
func UsersUpdate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
	request.ParseForm()
	fmt.Printf("put request %+v", request.Form)
	println("UsersUpdate")

	enabledFormFields := [1]string{"mail"}
	user := model.User{ID: routeParams.ByName("id")}
	for _, field := range enabledFormFields {
		// Get data from field
		value := request.Form.Get(field)
		if value == "" {
			continue
		}
		err := validateFormField(value, field)
		if err != nil {
			JSONResponse := service.JSONError{
				Status:        "Error",
				HTTPErrorCode: http.StatusBadRequest,
				Message:       "Invalid form value in request: " + value + " for field: " + field,
			}
			JSONResponse.Send(response)
			return
		}
		// AddValidFieldValuentf("mail after validation: %+v", value)
		err = AddValidFieldValueToUser(field, value, &user)
		if err != nil {
			fmt.Printf("unkonwn err: %+v", err)
			JSONResponse := service.JSONError{
				Status:        "Error",
				HTTPErrorCode: http.StatusInternalServerError,
				Message:       "Unknown error when setting user's field: " + field,
			}
			JSONResponse.Send(response)
			return
		}
	}

	err := user.Update()
	if err != nil {
		JSONResponse := service.JSONError{
			Status:        "Error",
			HTTPErrorCode: http.StatusInternalServerError,
			Message:       "Eerror when saving Users",
		}
		fmt.Printf("Error: %+v", err)
		JSONResponse.Send(response)
		return
	}
	service.SendJSONResponse(user, response)
	return
}

func validateFormField(fieldValue string, fieldName string) (err error) {
	// Validate data
	validator := service.Validation{Value: fieldValue, FieldName: fieldName}
	validator.Run()
	if !validator.Valid {
		return
	}
	return
}

// AddValidFieldValueToUser changes a field's value in the Users struct
func AddValidFieldValueToUser(field string, value string, user *model.User) (err error) {
	// Set field name to Uppercase
	field = strings.Title(field)
	// Fetch the field reflect.Value
	structValue := reflect.ValueOf(user).Elem()
	structFieldValue := structValue.FieldByName(field)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", field)
	}

	// If obj field value is not settable an error is thrown
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", value)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)

	return nil
}
