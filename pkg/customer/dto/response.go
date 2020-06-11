package dto

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
)

//swagger:model customer-response-dto-model
type CustomerResponseDto struct {
	// The ResponseData status
	//
	defaultresponse.ResponseData
	Data domain.Customer `json:"data"`
}

func NewCustomerResponseDto(status bool, message string, customer domain.Customer) *CustomerResponseDto {
	response := CustomerResponseDto{}
	response.Status = status
	response.Message = message
	response.Data = customer
	return &response
}
