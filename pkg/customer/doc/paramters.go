package doc

import "github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"

// swagger:parameters createCustomer
type CreateUserRequestWrapper struct {
	// in: body
	Body struct {
		dto.CustomerRequest
	}
}



