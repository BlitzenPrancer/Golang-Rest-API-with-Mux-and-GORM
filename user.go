package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database variable
var DB *gorm.DB

// error variable
var err error

// constant variable for database url
const DNS = "root:suda@123@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

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
func InitialMigration() {
	// defining the properties to connect to the database.
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	// checking the errors
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the Database")
	}
	// enabling auto migration to database
	DB.AutoMigrate(&User{}) // passing the reference to the struct
}

// creating handler methods to recieve all thedata from REST API's and save to the database.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting list of users available
	var users []User
	// fetching
	DB.Find(&users)
	// once after I get all the data I just have to encode and send back to client.
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Ill be getting which particular id that I want to fetch data
	params := mux.Vars(r)
	var user User
	// finding data in databse
	DB.First(&user, params["id"])
	// passing the data back
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// for http.ResponseWriter I'll be setting Content-Type to application/json
	w.Header().Set("Content-Type", "application/json")
	var user User
	// using json module to decode the data getting from the request body itself.
	json.NewDecoder(r.Body).Decode(&user) // all the data will be converted to user struct.
	// saving the data
	DB.Create(&user)
	// passing the data back to the browser
	json.NewEncoder(w).Encode(user) // encoding the data to ResponseWriter
}

func UpateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Ill be getting which particular id that I want to fetch data
	params := mux.Vars(r)
	var user User
	// getting value from database
	DB.First(&user, params["id"])
	// whatever data I'm getting from the body I'll just decode the information and update the reference of the same user.
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	// passing the data back to browser itself
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//setting the header information
	w.Header().Set("Content-Type", "application/json")
	// getting the parameter
	params := mux.Vars(r)
	var user User
	// getting the user and deleting user reference and param id
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user is deleted succesfully")
}
