package chttp

import (
	"encoding/json"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"net/http"
)

type customerControllerHandler struct {
	userService iservice.ICustomerService
}

//CreateUser calls the IUserService which is implemented by UserService
func (u customerControllerHandler) create(w http.ResponseWriter, r *http.Request) {
	var request dto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := defaultresponse.NewResponseDto(false, "Error while decoding request body",nil)
		json.NewEncoder(w).Encode(response)
		return
		//return
	}

	//response := u.userService.CreateUser(user)
	//utils.Respond(w, response)
}
