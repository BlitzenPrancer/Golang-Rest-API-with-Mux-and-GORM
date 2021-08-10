package main

import (
	"net/http"

	"gorm.io/gorm"
)

// database variable
var DB *gorm.DB

// error variable
var err error

// constant variable for database url
const DNS = "root:@suda@123(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

// I am defining GORM over here and also define the struct over here
// this struct will help me to store data into my database.
// once defined the struct ill be sending and recieving data through json when working with the Rest API's
type User struct {
	gorm.Model // converting the struct into orm model so that I can use this model to save the data and get the data from the database
	// defing what those properties will be when sending and recieving data in json
	FirstName string `json:"firstname"` // when passing data from json response body ill be giving first name over here.
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

// function to create the database and enable auto-migration
func initialMigration() {
	DB, err = gorm.Open()
}

// creating handler methods
func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
