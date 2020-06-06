package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/uuid"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"log"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(
	repository irepository.ICustomerRepository,
	userRepository irepository.IUserRepository,
	tokenService iservice.ITokenService,
	mailService mail.IMail, uuid uuid.IUuid) iservice.ICustomerService {
	return customerService{
		repository,
		userRepository,
		tokenService,
		mailService,
		uuid,
	}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository     irepository.IUserRepository
	tokenService  	iservice.ITokenService
	mailService        mail.IMail
	uuid uuid.IUuid
}

//refactor and test case needed
//Validate and crease customer
func (s customerService) CreateUser(request dto.CustomerRequest) (*domain.Customer, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	if found := s.userRepository.FindUserExist(request.Email); found {
		return nil, errors.New("user Already Exist with email")
	}

	customer := domain.Customer{}
	customer.FirstName = request.FirstName
	customer.LastName = request.LastName
	customer.Email = request.Email
	customer.Address = request.Address
	customer.PhoneNumber = request.PhoneNumber
	customer.Password = request.Password
	customer.Age = request.Age

	newCustomer, err := s.customerRepository.Save(customer)
	if err != nil {
		return nil, err
	}
	//
	result, err := s.sendCustomerEmail(*newCustomer)
	//
	if err != nil {
		return nil, err
	}
	log.Print(result)
	return newCustomer, nil
}

func (s customerService) sendCustomerEmail(customer domain.Customer) (string, error) {
	uniqueId, linkToSend := s.generateLinkToSendToUser()

	_, err := s.tokenService.CreateTokenWithExpirationInHours(customer.UserId, uniqueId,1)
	if err != nil {
		return "", err
	}
	email := s.createMail(customer, linkToSend)

	feedback, err := s.mailService.SendEMail(*email)

	if err != nil {
		return "", errors.New("error occurred try to send mail, try again later")
	}

	return "Mail sent successfully" + feedback, nil

}


func (s customerService) createMail(customer domain.Customer, linkToSend string) *mail.EMailMessage {
	subject := "DkreativeCoders Verify User"
	text := "Please visit this link to verify your account. \n This links expires in an hour \n" + linkToSend
	recipient := customer.Email
	email := mail.NewEMailMessage(subject, text, recipient, nil)
	return email
}

func (s customerService) generateLinkToSendToUser() (string, string) {
	uniqueId := s.uuid.GenerateUniqueId()
	linkToSend := "http/Dmessanger:8900/verify-user/" + uniqueId
	return uniqueId, linkToSend
}
