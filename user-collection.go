package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
)

const (
  configFile = "collections/users.json"
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

func getUserByID(id string) (user User, err error) {
  // Open File
  users, err := readUsersFromFile()
  if err != nil {
    return user, err
  }
  for _, user := range users {
    if user.Id == id {
      return user, nil
    }
  }

  err = fmt.Errorf("No matching user found by Id %s", id)
  // No match?
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