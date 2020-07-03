package doc

import "github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"

//The registration request body
// swagger:parameters createCustomer
type CreateUserRequestWrapper struct {
	// in: body
	Body dto.CustomerRequest
}

// The activate user url parameter
// swaggr:parameters activateCustomer
//type ActivateUserRequestWrapper struct {
//	// A long sequence of character
//	// in: query
//	token string
//
//	//A string of 6 digit
//	// in: query
//	// required: true
//	Otp string
//}
