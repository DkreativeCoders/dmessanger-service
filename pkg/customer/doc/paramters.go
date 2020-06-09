package doc

import "github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"

//The registration request body
// swagger:parameters createCustomer
type CreateUserRequestWrapper struct {
	// in: body
	Body dto.CustomerRequest
}

// The activate user url parameter
// swagger:parameters activateCustomer
type ActivateUserRequestWrapper struct {
	// in: path
	Token int
}
