package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/otp"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/uuid"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"log"
	"time"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(
	repository irepository.ICustomerRepository,
	userRepository irepository.IUserRepository,
	tokenRepository irepository.ITokenRepository,
	tokenService iservice.ITokenService,
	mailService mail.IMail,
	uuid uuid.IUuid,
	otp otp.IOtp,
	) iservice.ICustomerService {
	return customerService{
		repository,
		userRepository,
		tokenRepository,
		tokenService,
		mailService,
		uuid,
		otp,
	}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository     irepository.IUserRepository
	tokenRepository    irepository.ITokenRepository
	tokenService       iservice.ITokenService
	mailService        mail.IMail
	uuid               uuid.IUuid
	otp				   otp.IOtp
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

func (s customerService) ActivateUser(tk string) error {
	// Check for token in the repository
	token, err := s.tokenRepository.FindByToken(tk)
	if err != nil {
		return err
	}

	// Check if token has expired
	tokenExpiryTime := token.ExpiresOn
	if time.Now().After(tokenExpiryTime) {
		return errors.New("token expired")
	}

	// Find user with token
	user, er := s.userRepository.FindByID(int(token.UserId))
	if er != nil {
		return er
	}

	// Activate user
	user.IsActive = true
	s.userRepository.Update(*user)

	return nil
}

func (s customerService) sendCustomerEmail(customer domain.Customer) (string, error) {

	uniqueId, linkToSend := s.generateLinkToSendToUser()
	otp := s.otp.GenerateOTP()

	_, err := s.tokenService.CreateTokenWithExpirationInHours(customer.UserId, uniqueId, otp, 1)
	if err != nil {
		return "", err
	}
	email := s.createMail(customer, linkToSend, otp)

	feedback, err := s.mailService.SendEMail(*email)

	if err != nil {
		return "", errors.New("error occurred try to send mail, try again later")
	}

	return "Mail sent successfully" + feedback, nil

}

func (s customerService) createMail(customer domain.Customer, linkToSend, otp string) *mail.EMailMessage {
	subject := "DkreativeCoders Verify User"
	text := "Please visit this link to verify your account. \n This links expires in an hour \n" + linkToSend +
		"\n You can also use this OTP to verify your account via your mobile Device " + otp
	recipient := customer.Email
	email := mail.NewEMailMessage(subject, text, recipient, nil)
	return email
}

func (s customerService) generateLinkToSendToUser() (string, string) {
	uniqueId := s.uuid.GenerateUniqueId()
	linkToSend := "https://dmessanger-service.herokuapp.com/verify-user/" + uniqueId

	return uniqueId, linkToSend
}
