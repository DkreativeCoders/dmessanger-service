package chttp

import (
	"encoding/json"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/gorilla/mux"
	"net/http"
)


func NewCustomerHandler(router *mux.Router, customerService iservice.ICustomerService) {
	handler := &customerControllerHandler{
		customerService: customerService,
	}

	router.HandleFunc("/api/v1/customers", handler.create).Methods("POST")

	//return userControllerHandler{userService}
}

type customerControllerHandler struct {
	customerService iservice.ICustomerService
}



//swagger definition
//Write test case for it
//create extract request to
func (c customerControllerHandler) create(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /api/v1/customers createCustomer
	//
	// Creates a new customer
	// ---
	// Consumes:
	//	- application/json
	// Produces:
	//  - application/json
	// Responses:
	//   default:
	//     "$ref": "#/responses/customerCreatedResponse"
	//   200:
	//     "$ref": "#/responses/customerCreatedResponse"
	//   400:
	//     "$ref": "#/responses/customerBadRequestResponse"
	//   422:
	//     "$ref": "#/responses/customerErrorResponse"


	var request dto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errResponse := defaultresponse.NewResponseDto(false, "Error while decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
		//return
	}
	customer, errorRes := c.customerService.CreateUser(request)

	if errorRes != nil{
		errResponse := defaultresponse.NewResponseDto(false, errorRes.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	customer.UserId=6
	response := defaultresponse.NewResponseDto(true, "Successful")

	json.NewEncoder(w).Encode(response)

}
