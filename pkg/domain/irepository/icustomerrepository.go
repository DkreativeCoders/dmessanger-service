package irepository

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// IUserRepository
// Define your basic crud operation in this interface

type ICustomerRepository interface {
	FindByID(id int) (*domain.Customer, error)
	FindByUserId(customerId int) (*domain.Customer, error)
	FindByEmail(email string) (*domain.Customer, error)
	FindAll() []domain.Customer
	Save(customer domain.Customer) (*domain.Customer, error)
	Update(customer domain.Customer) (*domain.Customer, error)
}
