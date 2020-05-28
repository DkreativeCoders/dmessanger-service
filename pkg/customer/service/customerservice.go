package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(repository irepository.ICustomerRepository,
	userRepository irepository.IUserRepository,
	tokenRepository irepository.ITokenRepository,
	mailService mail.IMail ) iservice.ICustomerService{
	return customerService{
		repository,
		userRepository,
		tokenRepository,
		mailService}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository irepository.IUserRepository
	tokenRepository irepository.ITokenRepository
	mailService mail.IMail
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
	result,err := s.sendCustomerEmail(*newCustomer)
	//
	if err != nil {
		return nil, err
	}
	log.Print(result)
	return newCustomer,nil
}

func (s customerService) sendCustomerEmail(customer domain.Customer) (string, error){
	uniqueId, linkToSend := s.generateLinkToSendToUser()

	token := domain.Token{}
	token.UserId=customer.UserId
	token.Token=uniqueId
	token.ExpiresOn=time.Now().Add(1 * time.Hour)

	_, err := s.tokenRepository.Create(token)
	if err !=nil{
		return "",errors.New("error occurred, try again")
	}

	subject:="Verify User"
	text :="Please visit this link to very your account. \n This links expires in an hour \n"+ linkToSend
	recipient:=customer.Email
	feedback, err := s.mailService.SendMail(subject, text, recipient)
	if err !=nil{
		return "",errors.New("error occurred try to send mail, try again later")
	}


	return "Mail sent successfully"+feedback, nil
}

func (s customerService) generateLinkToSendToUser() (string, string){
	uniqueId := uuid.NewV4().String()
	linkToSend :="http/Dmessanger:8900/verify-user/"+uniqueId
	return uniqueId,linkToSend
}

