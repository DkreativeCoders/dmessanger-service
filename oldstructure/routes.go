package oldstructure

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/oldstructure/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Server() *http.Server {

	router := mux.NewRouter()

	router.HandleFunc("/api/create/user", controllers.UserController.CreateUser).Methods("POST")
	router.HandleFunc("/api/get/users", controllers.UserController.GetAllUser).Methods("GET")
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	srv := &http.Server{Handler: router, Addr: ":"+port,}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	return srv
}
