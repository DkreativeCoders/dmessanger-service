package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danieloluwadare/dmessanger/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//GetAllUser
	router.HandleFunc("/api/create/user", controllers.UserController.CreateUser).Methods("POST")
	router.HandleFunc("/api/get/users", controllers.UserController.GetAllUser).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
