package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/* Steps
1) Create routers
2) Initialize routers with mux
3) Initialize database connection with the database using GORM
*/

// Method that will initialize the router
func initializeRouter() {
	// creating the router
	r := mux.NewRouter()
	// defining different API's that I want to use.
	// Here I'll take the example of an user where I'll store the data of the user to my database.
	// Implemeting different routes over here
	r.HandleFunc("/users", GetUsers).Methods("GET")
	// getting user along with id
	r.HandleFunc("/users{id}", GetUser).Methods("GET")
	// create user
	r.HandleFunc("/users", CreateUser).Methods("POST")
	// delete user
	r.HandleFunc("/users{id}", DeleteUser).Methods("DELETE")
	// update user
	r.HandleFunc("/users{id}", UpateUser).Methods("PUT")

	// now i have to start the server and serve at port 9000
	log.fatal(http.ListenAndServe(":9000", r))

}

func main() {
	initializeRouter()
}