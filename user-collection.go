package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/satori/go.uuid"
)

const (
  configFile = "collections/users.json"
  idLength = 8
)


func init() {
  // get config

  connectToMongo()
}

var mongoSession *mgo.Session
var userCollection *mgo.Collection
type mongoResult interface {} 

func connectToMongo() {
  var err error
  mongoSession, err = mgo.Dial("localhost:27017")
  if err != nil {
    panic("Can't connect to database")
  }
  mongoSession.SetMode(mgo.Monotonic, true)
  userCollection = mongoSession.DB("testgo").C("users")
}

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
  ID string
  Mail string
  Accounts []account
  Posts []post
}

// generateId creates a pseudo ranodom string for the user Id 
func (userStruct *User) generateID() {
  uuidID := []byte(uuid.NewV1().String())
  // Cut to 8 char length 
  userStruct.ID = string(uuidID[:idLength])
}

func (userStruct *User) insert() (err error) {
  err = userCollection.Insert(&userStruct)
  return err
}

// Find public method of User struct to find a user by id
func (userStruct *User) Find(id string) (user User, err error) { 
  user, err = getUserByID(id)
  if err != nil {
    // no user found
    return user, err
  }
  
  return user, nil
}

// FindAll finds all users from collections
func (userStruct *User) FindAll() (users []User, err error) {
  query := bson.M{}
  fmt.Printf("mongo query %+v", query)
  err = userCollection.Find(query).All(&users)
  if err != nil {
    return users, err
  }
  return users, nil
}

func getUserByID(id string) (user User, err error) {
  return findBy("id", id)
}

func getUserByMail(mail string) (user User, err error) {
  return findBy("mail", mail)
}

// Find user entity by some of its value
// field: name of field, value: value to search for the field
func findBy(field string, value string) (user User, err error) {
  // query := bson.M{field: value}
  query := bson.M{field: value}
  fmt.Printf("mongo query %+v", query)
  err = userCollection.Find(query).One(&user)
  fmt.Printf("mongo result %+v", user)
  if err != nil {
    return user, err
  }
  
  return user, nil
}