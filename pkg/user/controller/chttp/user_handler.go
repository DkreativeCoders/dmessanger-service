package chttp

import (
	"encoding/json"
	"github.com/DkreativeCoders/dmessanger-service/pkg/constanst"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewUserHandler(router *mux.Router, userService iservice.IUserService) {
	handler := &userControllerHandler{
		userService: userService,
	}

	router.HandleFunc(constanst.ApiVersion1+"users", handler.create).Methods("POST")
	router.HandleFunc(constanst.ApiVersion1+"users", handler.getAll).Methods("GET")
	router.HandleFunc(constanst.ApiVersion1+"users/update-password/{userID}", handler.updatePassword).Methods("PATCH")
	router.HandleFunc(constanst.ApiVersion1+"users/enable-user/{userID}", handler.enableUser).Methods("PATCH")
	router.HandleFunc(constanst.ApiVersion1+"users/disable-user/{userID}", handler.disableUser).Methods("PATCH")
	router.HandleFunc(constanst.ApiVersion1+"login", handler.authenticateUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("users/forgot-password/{email}", handler.forgotPassword).Methods("POST")
	router.HandleFunc("users/reset-password/{token}", handler.resetPassword).Methods("POST")

	//return userControllerHandler{userService}

}

type userControllerHandler struct {
	userService iservice.IUserService
}

func (u userControllerHandler) authenticateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /api/v1/authenticate userAuthentication
	//
	// All User Authentication
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"
	//   200:
	//     "$ref": "#/responses/tokenResponse"
	//   400:
	//     "$ref": "#/responses/badRequestResponse"
	//   401:
	//     "$ref": "#/responses/unAuthenticatedResponse"

	//setupResponse(&w, r)
	//if r.Method == "OPTIONS" {
	//	w.WriteHeader(http.StatusOK)
	//	return
	//}
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errResponse := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	tokenResponse, err := u.userService.Login(request)

	if err != nil {
		errResponse := defaultresponse.NewResponseDto(false, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	json.NewEncoder(w).Encode(tokenResponse)

	//customer, errorRes := c.customerService.CreateUser(request)

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

func (u userControllerHandler) enableUser(w http.ResponseWriter, r *http.Request) {

	// swagger:operation PUT /api/v1/users/enable-user/{UserID} enableUser
	//
	// Sets the isEnabled field of a user's to true
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

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

	if err != nil {
		response := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		json.NewEncoder(w).Encode(response)
		return
	}

	serviceError := u.userService.EnableUser(userID)
	if serviceError != nil {
		response = defaultresponse.NewResponseDto(false, serviceError.Error())
	} else {
		response = defaultresponse.NewResponseDto(true, "Successful")
	}

	json.NewEncoder(w).Encode(response)
}

func (u userControllerHandler) disableUser(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /api/v1/users/disable-user/{UserID} disableUser
	//
	// Sets the isEnabled field of a user's to false
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

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

	if err != nil {
		response := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		json.NewEncoder(w).Encode(response)
		return
	}

	serviceError := u.userService.DisableUser(userID)
	if serviceError != nil {
		response = defaultresponse.NewResponseDto(false, serviceError.Error())
	} else {
		response = defaultresponse.NewResponseDto(true, "Successful")
	}

	json.NewEncoder(w).Encode(response)
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

func (u userControllerHandler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /api/v1/users/forgot-password/{Email} forgotPassword
	//
	// Initate process to reset a user's password
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

	vars := mux.Vars(r)
	email := vars["email"]
	w.Header().Add("Content-Type", "application/json")

	err := u.userService.ForgotPassword(email)
	var response *defaultresponse.ResponseData

	if err != nil {
		response = defaultresponse.NewResponseDto(false, err.Error())
	} else {
		response = defaultresponse.NewResponseDto(true, "An email has been sent to the provided email address with instructions on how to reset your password")
	}

	json.NewEncoder(w).Encode(response)

}

func (u userControllerHandler) resetPassword(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /api/v1/users/reset-password/{Token} resetPassword
	//
	// Reset a user's password
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

	vars := mux.Vars(r)
	token := vars["token"]
	w.Header().Add("Content-Type", "application/json")

	var resetPasswordRequest dto.ResetPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&resetPasswordRequest)

	if err != nil {
		response := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		json.NewEncoder(w).Encode(response)
		return
	}

	err = u.userService.ResetPassword(token, resetPasswordRequest)

	var response *defaultresponse.ResponseData

	if err != nil {
		response = defaultresponse.NewResponseDto(false, err.Error())
	} else {
		response = defaultresponse.NewResponseDto(true, "Password reset successfully")
	}

	json.NewEncoder(w).Encode(response)

}
