package chttp

import (
	"encoding/json"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewUserHandler(router *mux.Router, userService iservice.IUserService) {
	handler := &userControllerHandler{
		userService: userService,
	}

	router.HandleFunc("/api/v1/users", handler.create).Methods("POST")
	router.HandleFunc("/api/v1/users", handler.getAll).Methods("GET")
	router.HandleFunc("/api/v1/users/update-password/{userID}", handler.updatePassword).Methods("PATCH")

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
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		//return
	}
	//response := u.userService.CreateUser(user)
	//utils.Respond(w, response)
}

//GetAllUser This is the method called from the route to fetch all user from the service class
func (u userControllerHandler) getAll(w http.ResponseWriter, r *http.Request) {
	response := u.userService.GetAllUser()
	utils.Respond(w, response)
}

func (u userControllerHandler) getUser(w http.ResponseWriter, r *http.Request) {

}

func (u userControllerHandler) updatePassword(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /api/v1/users/update-password/{UserID} updatePassword
	//
	// Updates a user's password
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

	var updatePasswordRequest dto.UpdatePasswordRequest
	vars := mux.Vars(r)
	userIDVar := vars["userID"]
	userID, err := strconv.Atoi(userIDVar)
	w.Header().Add("Content-Type", "application/json")

	var response *defaultresponse.ResponseData

	if err != nil {
		response = defaultresponse.NewResponseDto(false, "User Id must be an integer")
		json.NewEncoder(w).Encode(response)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updatePasswordRequest)

	if err != nil {
		response := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		json.NewEncoder(w).Encode(response)
		return
	}

	serviceError := u.userService.UpdatePassword(userID, updatePasswordRequest)
	if serviceError != nil {
		response = defaultresponse.NewResponseDto(false, serviceError.Error())
	} else {
		response = defaultresponse.NewResponseDto(true, "Successful")
	}

	json.NewEncoder(w).Encode(response)

}
