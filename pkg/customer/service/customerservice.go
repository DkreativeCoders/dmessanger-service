package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(repository irepository.ICustomerRepository,userRepository irepository.IUserRepository ) iservice.ICustomerService{
	return customerService{repository, userRepository}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository irepository.IUserRepository
	//userService iservice.IUserService

}
//refactor and test case needed
//Validate and crease customer
func (s customerService) CreateUser(request dto.CustomerRequest) (*domain.Customer, error){
	 err := request.Validate()
	if err != nil {
		return nil, err
	}

	//user := domain.User{
	//	FirstName: request.FirstName,
	//	LastName: request.LastName,
	//	Age: request.Age,
	//	Email: request.Email,
	//	PhoneNumber: request.PhoneNumber,
	//	Password: request.Password,
	//	Address: request.Address,
	//}
	//
	//if err := user.ValidateToError(); err!=nil {
	//	return nil, err
	//}

	if found := s.userRepository.FindUserExist(request.Email); found{
		return nil,errors.New("user Already Exist with email")
	}

	//newUser, err := s.userService.CreateUser(user)

	//if err != nil {
	//	return nil, err
	//}

	customer := domain.Customer{}
	customer.FirstName=request.FirstName
	customer.LastName=request.LastName
	customer.Email=request.Email
	customer.Address=request.Address
	customer.PhoneNumber=request.PhoneNumber
	customer.Password=request.Password
	customer.Age=request.Age

	//customer := domain.Customer{
	//	UserId: newUser.ID,
	//	User: user,
	//}

	newCustomer, err :=s.customerRepository.Save(customer)
	if err != nil {
		return nil, err
	}
	return newCustomer,nil
}
