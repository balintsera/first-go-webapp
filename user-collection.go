package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "github.com/satori/go.uuid"
  "reflect"
)

const (
  configFile = "collections/users.json"
  idLength = 8
)

// @TODO move this to a dedicated dir ?
type account struct {
  id string
  title string
  url string
  oauthToken string
}

type post struct {
  id string
  content string
  url string
  date string
}

// User base class. Has accounts and has posts
type User struct {
  Id string
  Mail string
  accounts []account
  posts []post
}

// generateId creates a pseudo ranodom string for the user Id 
func (userStruct *User) generateId() {
  uuidId := []byte(uuid.NewV1().String())
  // Cut to 8 char length 
  userStruct.Id = string(uuidId[:idLength])
}

// Find public method of User struct to find a user by id
func (userStruct *User) Find(id string) (user User, err error) {
  userFound, err := getUserByID(id)
  if err != nil {
    // no user found
    return user, err
  }

  userStruct.Id = userFound.Id
  userStruct.Mail = userFound.Mail
  return userFound, nil
}

// FindAll finds all users from collections
func (userStruct *User) FindAll() (users []User, err error) {
  users, err = readUsersFromFile()
  if err != nil {
    return users, err
  }
  return users, nil
}

func getUserByID(id string) (user User, err error) {
  return findBy("Id", id)
}

func getUserByMail(mail string) (user User, err error) {
  return findBy("Mail", mail)
}

// Find user entity by some of its value
// field: name of field, value: value to search for the field
func findBy(field string, value string) (user User, err error) {
  // Open File
  users, err := readUsersFromFile()
  if err != nil {
    return user, err
  }
  // Walk trough all the users, return when value found in field
  var fieldValue string
  for _, user := range users {
    // Get the value of the field. Works only when the field's value type is a string
    fieldValue = reflect.ValueOf(user).FieldByName(field).String()
    // Found?
    if fieldValue == value {
      return user, nil
    }
  }

  // Not found
  err = fmt.Errorf("No matching user found by %s: %s", field, value)
  return user, err
}

func readUsersFromFile() (users []User, err error) {
  file, err := ioutil.ReadFile(configFile)
  if err != nil {
    return users, err 
  }

  // Luca 14:45-kor aludt el
  
  err = json.Unmarshal(file, &users)
  if err != nil {
    fmt.Println("error: ", err)
    return users, err
  }
  
  return users, nil
}