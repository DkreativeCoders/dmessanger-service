package controllers

import (
	"encoding/json"
	util "github.com/danieloluwadare/dmessanger/utils"
	"net/http"

	"github.com/danieloluwadare/dmessanger/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/danieloluwadare/dmessanger/repository"
	"github.com/danieloluwadare/dmessanger/service"
)

var UserController UserControllerHandler

func init() {

	db := models.GetDB()
	userRepository := repository.NewInMemoryRepository(db)
	userService := service.NewService(userRepository)
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

func (u UserControllerHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

//func createEvent(w http.ResponseWriter, r *http.Request) {
//	var newEvent event
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
//	}
//
//	json.Unmarshal(reqBody, &newEvent)
//	events = append(events, newEvent)
//	w.WriteHeader(http.StatusCreated)
//
//	json.NewEncoder(w).Encode(newEvent)
//}

