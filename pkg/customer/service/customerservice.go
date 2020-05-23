package service

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(repository irepository.ICustomerRepository,userService iservice.IUserService ) iservice.ICustomerService{
	return cutomerservice{repository, userService}
}

type cutomerservice struct {
	repository irepository.ICustomerRepository
	userService iservice.IUserService

}
//refactor and test case needed
//Validate and crease customer
func (s cutomerservice) CreateUser(request dto.CustomerRequest) (*domain.Customer, error){
	 err := request.Validate()
	if err != nil {
		return nil, err
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName: request.LastName,
		Age: request.Age,
		Email: request.Age,
		PhoneNumber: request.PhoneNumber,
		Password: request.Password,
		Address: request.Address,
	}

	if err := user.ValidateToError(); err!=nil {
		return nil, err
	}

	newUser, err := s.userService.CreateUser(user)

	if err != nil {
		return nil, err
	}

	customer := domain.Customer{
		UserId: newUser.ID,
		User: user,
	}

	newCustomer, err :=s.repository.Save(customer)
	if err != nil {
		return nil, err
	}
	return newCustomer,nil
}
