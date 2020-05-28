package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/satori/go.uuid"

)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(repository irepository.ICustomerRepository,
	userRepository irepository.IUserRepository,
	mailService mail.IMail ) iservice.ICustomerService{
	return customerService{repository, userRepository,mailService}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository irepository.IUserRepository
	mailService mail.IMail
	//userService iservice.IUserService

}
//refactor and test case needed
//Validate and crease customer
func (s customerService) CreateUser(request dto.CustomerRequest) (*domain.Customer, error){
	 err := request.Validate()
	if err != nil {
		return nil, err
	}

	if found := s.userRepository.FindUserExist(request.Email); found{
		return nil,errors.New("user Already Exist with email")
	}

	customer := domain.Customer{}
	customer.FirstName=request.FirstName
	customer.LastName=request.LastName
	customer.Email=request.Email
	customer.Address=request.Address
	customer.PhoneNumber=request.PhoneNumber
	customer.Password=request.Password
	customer.Age=request.Age

	newCustomer, err :=s.customerRepository.Save(customer)
	if err != nil {
		return nil, err
	}
	//
	_, _ = s.sendCustomerEmail(*newCustomer)
	//
	return newCustomer,nil
}

func (s customerService) sendCustomerEmail(customer domain.Customer) (string, error){

	//s.mailService.SendMail()
	_, _ = s.generateLinkToSendToUser()
//implement your email sending here @AB
	return "nil", nil
}

func (s customerService) generateLinkToSendToUser() (string, string){
	uniqueId := uuid.NewV4().String()
	linkToSend :="/verify-user/"+uniqueId
	return uniqueId,linkToSend
}

