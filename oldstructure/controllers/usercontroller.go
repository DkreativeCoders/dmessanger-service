package controllers

import (
	"encoding/json"
	"github.com/danieloluwadare/dmessanger/oldstructure/migrations"
	util "github.com/danieloluwadare/dmessanger/oldstructure/utils"
	"net/http"

	"github.com/danieloluwadare/dmessanger/oldstructure/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/oldstructure/models"
	"github.com/danieloluwadare/dmessanger/oldstructure/repository"
	"github.com/danieloluwadare/dmessanger/oldstructure/service"
)

var UserController UserControllerHandler

func init() {

	db := migrations.GetDB()
	userRepository := repository.INewUserRepository(db)
	userService := service.INewService(userRepository)
	usehandler := NewUserHandler(userService)
	UserController = usehandler
}

func NewUserHandler(userService iservice.IUserService) UserControllerHandler {
	return UserControllerHandler{userService}
}

type UserControllerHandler struct {
	userService iservice.IUserService
}

//CreateUser calls the IUserService which is implemented by UserService
func (u UserControllerHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.Respond(w, util.Message(false, "Error while decoding request body"))
		//return
	}
	response := u.userService.CreateUser(user)
	util.Respond(w, response)
}

//GetAllUser This is the method called from the route to fetch all user from the service class
func (u UserControllerHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	response := u.userService.GetAllUser()
	util.Respond(w, response)
}

func (u UserControllerHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}



