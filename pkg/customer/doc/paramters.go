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
	// A long sequence of character
	// in: query
	// name: token
	Token string

	//A string of 6 digit
	// in: query
	// name: otp
	// required: true
	// example: 000111
	Otp string
}
