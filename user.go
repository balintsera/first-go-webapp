package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
  "fmt"
  "reflect"
  "log"
)

type jsonResponseSourceObject interface {}

// UserIndex lists users
func UserIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserIndex")
  var user User
  users, err := user.FindAll()
  if err != nil {
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusNotFound, 
      Message: "No user found",
    }
    JSONResponse.Send(response)
    return
  } 

  // Send user list
   sendJSONResponse(users, response)
}

// UserCreate create a new user via POST
func UserCreate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserCreate")
  
  request.ParseForm()
  
  // Validation: is mail field set?
  if len(request.Form.Get("mail")) < 1 {
    // Not set, send error
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusBadRequest, 
      Message: "'mail' parameter is mandatory",
    }
    JSONResponse.Send(response)
    return
  }

  // Validate email via regexp
  validator := Validation{fieldType: FieldTypeMail, value: request.Form.Get("mail")}
  validator.Run()
  if !validator.valid {
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusBadRequest, 
      Message: "'mail' parameter is not valid. Please use a valid mail address. You sent: " + request.Form.Get("mail"),
    }
    JSONResponse.Send(response)
    return
  }

  // Check for existence
  found, _ := getUserByMail(request.Form.Get("mail"))
  fmt.Printf("found mail: %+v", found)
  // err is true if not found
  if len(found.Mail) > 0 {
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusBadRequest, 
      Message: "A user is already registered with this mail",
    }
    JSONResponse.Send(response)
    return
  }
  
  fmt.Printf("mail %+v", request.Form.Get("mail"))

  // Creation
  user := User{Mail: request.Form.Get("mail")}
  user.generateID()
  user.insert()
  
  // @TODO insert to database
  fmt.Printf("new user: %+v", user)
  sendJSONResponse(user, response)
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
  request.ParseForm()
  fmt.Printf("put request %+v", request.Form)
  println("UserUpdate")
 
  enabledFormFields := [1]string{"mail"}
  user := User{ID: routeParams.ByName("id")}
  for _, field := range enabledFormFields {
    // Get data from field
    value := request.Form.Get(field)
    if value == "" {
      continue
    }
    err := validateFormField(value, field)
    if err != nil {
        JSONResponse := JSONError {
          Status: "Error",
          HTTPErrorCode: http.StatusBadRequest, 
          Message: "Invalid form value in request: " + value + " for field: " + field,
        }
        JSONResponse.Send(response)
        return
    }
    // AddValidFieldValuentf("mail after validation: %+v", value)
    err = AddValidFieldValueToUser(field, value, &user)
    if err != nil {
      JSONResponse := JSONError {
          Status: "Error",
          HTTPErrorCode: http.StatusInternalServerError, 
          Message: "Unknown error when setting user field: " + field,
        }
        JSONResponse.Send(response)
        return
    }

  }
  err := user.Update()
  if err != nil {
    JSONResponse := JSONError {
      Status: "Error",
      HTTPErrorCode: http.StatusInternalServerError, 
      Message: "Eerror when saving user",
    }
    fmt.Printf("Error: %+v", err)
    JSONResponse.Send(response)
    return
  }
}

func validateFormField(fieldValue string, fieldName string) (err error) {
  // Validate data
  validator := Validation{value: fieldValue, fieldName: fieldName}
  validator.SetRule(0)
  validator.Run()
  if !validator.valid {
      return 
  }
  return
}

// AddValidFieldValue changes a field's value in the user struct
func AddValidFieldValueToUser(field string, value string, user *User) (err error) {
  userStruct := reflect.ValueOf(&user)
  if userStruct.Elem().Kind() != reflect.Struct {
    return fmt.Errorf("Not a struct")
  }
  log.Printf("user after setting fields %+v", user)

  userField := userStruct.FieldByName(field)
  if !userField.IsValid() {
    return fmt.Errorf("No such field in struct: %s", field)
  }  
  userField.SetString(value)
  return nil
}

func sendJSONResponse(object jsonResponseSourceObject, response http.ResponseWriter) {
  response.Header().Set("Content-Type", "application/json")
  jsonResponse, err := json.Marshal(object)
  if err != nil {
    panic("JSON conversion failed")
  }
  response.Write(jsonResponse)
}
