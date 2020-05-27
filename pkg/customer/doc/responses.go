package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
)

// A Customer Created Response .
//swagger:response customerCreatedResponse
type CustomerResponseWrapper struct {
	// in: body
	Body dto.CustomerResponseDto

}

// A Customer Error Response .
//swagger:response customerErrorResponse
type CustomerErrorResponseWrapper struct {
	// in: body
	Body defaultresponse.ResponseData

}

// A Customer Bad Request Error Response .
//swagger:response customerBadRequestResponse
type CustomerBadRequestResponseWrapper struct {
	// in: body
	Body defaultresponse.ResponseData

}