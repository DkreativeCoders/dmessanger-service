package doc

import "github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"

//The registration request body
// swagger:parameters createCustomer
type CreateUserRequestWrapper struct {
	// in: body
	Body dto.CustomerRequest
}
