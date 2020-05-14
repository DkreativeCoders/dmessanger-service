package chttp

import (
	"encoding/json"
	"github.com/danieloluwadare/dmessanger/newstructure/domain"
	"github.com/danieloluwadare/dmessanger/newstructure/domain/iservice"
	util "github.com/danieloluwadare/dmessanger/oldstructure/utils"
	"github.com/gorilla/mux"
	"net/http"
)



func NewUserHandler(router *mux.Router, userService iservice.IUserService)  {
	handler := &userControllerHandler{
		userService: userService,
	}

	router.HandleFunc("/api/create/user", handler.create).Methods("POST")
	router.HandleFunc("/api/get/users", handler.getAll).Methods("GET")

	//return userControllerHandler{userService}
}

type userControllerHandler struct {
	userService iservice.IUserService
}

//CreateUser calls the IUserService which is implemented by UserService
func (u userControllerHandler) create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.Respond(w, util.Message(false, "Error while decoding request body"))
		//return
	}
	response := u.userService.CreateUser(user)
	util.Respond(w, response)
}

//GetAllUser This is the method called from the route to fetch all user from the service class
func (u userControllerHandler) getAll(w http.ResponseWriter, r *http.Request) {
	response := u.userService.GetAllUser()
	util.Respond(w, response)
}

func (u userControllerHandler) getUser(w http.ResponseWriter, r *http.Request) {

}



