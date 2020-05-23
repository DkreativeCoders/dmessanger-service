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

///create extract request to
func (c customerControllerHandler) create(w http.ResponseWriter, r *http.Request) {
	var request dto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errResponse := defaultresponse.NewResponseDto(false, "Error while decoding request body",nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
		//return
	}
	customer, errorRes := c.customerService.CreateUser(request)

	if errorRes != nil{
		errResponse := defaultresponse.NewResponseDto(false, "service",errorRes.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := defaultresponse.NewResponseDto(true, "Successful",customer)

	json.NewEncoder(w).Encode(response)

}
