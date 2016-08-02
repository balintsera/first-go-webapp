package main

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestAddValidFieldToUser(t *testing.T) {
    Convey("Given a user", t, func() {
        user := User{Mail: "balint@gmail.com", ID: "01234567"}
        Convey("When changing a valid field in user struct to 'anna@gmail.com'", func() {
            AddValidFieldValueToUser("Mail", "anna@gmail.com", &user)
            
            Convey("The value should be 'anna@gmail.com'", func() {
                So(user.Mail, ShouldEqual, "anna@gmail.com")

            })
        })
    })
}

