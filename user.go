package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
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
  id string
  mail string
  accounts []account
  posts []post
}

// UserIndex lists users
func UserIndex(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserIndex")
}

// UserCreate create a new user
func UserCreate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserCreate")
}

// UserShow display a user
func UserShow(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserShow")
}

// UserUpdate update a user
func UserUpdate(response http.ResponseWriter, request *http.Request, routeParams httprouter.Params) {
  println("UserUpdate")
}

