package iservice

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

type ICustomerService interface {
	CreateUser(request dto.CustomerRequest) (*domain.Customer, error)
	ActivateUser(tk string) error
}
