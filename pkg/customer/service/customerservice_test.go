package service_test

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/service"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//func TestCustomerService_CreateUser(t *testing.T) {
//
//	//	mock userRepository
//	//	mock customerRepository
//	//	mock tokenRepository
//	//	mock mailService
//
//	if testing.Short() {
//		t.Skip()
//	}
//
//	otpNumber := "000111"
//	timeAdded := time.Now().Add(1 * time.Hour)
//	mailtobesent := mail.NewEMailMessage("DkreativeCoders Verify User",
//		"Please visit this link to verify your account. \n This links expires in an hour \n"+"https://dmessanger-service.herokuapp.com/verify-user/unique-111\n You can also use this OTP to verify your account via your mobile Device "+otpNumber, "daniel@gmail.com",
//		nil)
//
//	testCases := []struct {
//		name    string
//		request dto.CustomerRequest
//
//		userRepoFindUserExistInput      string
//		userRepoFindUserExistReturnData bool
//
//		customerRepositorySaveInput        domain.Customer
//		customerRepositorySaveReturnOutput *domain.Customer
//		customerRepositorySaveReturnError  error
//
//		uniqueIdGenerated string
//
//		tokenServiceInputUserID               uint
//		tokenServiceInputUniqueID             string
//		tokenServiceInputExpirationTimeInHour time.Duration
//		tokenServiceCreateReturnOutPut        *domain.Token
//		tokenServiceCreateReturnError         error
//
//		mailServiceSendMailInput  mail.EMailMessage
//		mailServiceSendMailOutput string
//		mailServiceSendMailError  error
//
//		expectedValueDataResponse  *domain.Customer
//		expectedValueErrorResponse error
//	}{
//		{
//			"Test with valid customer request",
//			dto.CustomerRequest{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},
//
//			"daniel@gmail.com",
//			false,
//
//			domain.Customer{User: domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
//			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
//			nil,
//
//			"unique-111",
//
//			0,
//			"unique-111",
//			1,
//			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 0, Token: "unique-111", ExpiresOn: timeAdded},
//			nil,
//
//			*mailtobesent,
//			"success",
//			nil,
//
//			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
//			nil,
//		},
//
//		{
//			"Test customer already exist with already registered email",
//			dto.CustomerRequest{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},
//
//			"daniel@gmail.com",
//			true,
//
//			domain.Customer{User: domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
//			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
//			nil,
//
//			"unique-111",
//
//			0,
//			"unique-111",
//			1,
//			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 0, Token: "unique-111", ExpiresOn: timeAdded},
//			nil,
//
//			*mailtobesent,
//			"success",
//			nil,
//
//			nil,
//			errors.New("user Already Exist with email"),
//		},
//	}
//
//	for _, testCase := range testCases {
//
//		t.Run(testCase.name, func(t *testing.T) {
//			// Create dependency userRepo with mock implementation
//			userRepo := mocks.IUserRepository{}
//			userRepo.On("FindUserExist", testCase.userRepoFindUserExistInput).Return(testCase.userRepoFindUserExistReturnData)
//
//			// Create dependency customerRepo with mock implementation
//			customerRepo := mocks.ICustomerRepository{}
//			customerRepo.On("Save", testCase.customerRepositorySaveInput).Return(testCase.customerRepositorySaveReturnOutput, testCase.customerRepositorySaveReturnError)
//
//			// Create dependency tokenRepo with mock implementation
//			tokenService := mocks.ITokenService{}
//			tokenService.On("CreateTokenWithExpirationInHours", testCase.tokenServiceInputUserID, testCase.tokenServiceInputUniqueID, testCase.tokenServiceInputExpirationTimeInHour).Return(testCase.tokenServiceCreateReturnOutPut, testCase.tokenServiceCreateReturnError)
//
//			tokenRepo := mocks.ITokenRepository{}
//
//			// Create dependency tokenRepo with mock implementation
//			uuidService := mocks.IUuid{}
//			uuidService.On("GenerateUniqueId").Return(testCase.uniqueIdGenerated)
//
//			// Create dependency tokenRepo with mock implementation
//			mailService := mocks.IMail{}
//			mailService.On("SendEMail", testCase.mailServiceSendMailInput).Return(testCase.mailServiceSendMailOutput, testCase.mailServiceSendMailError)
//
//			//Todo: OTP service to be replaced with mock
//			otpService := mocks.IOtp{}
//			otpService.On("GenerateOTP").Return(otpNumber)
//
//			// Create userService and inject mock repo
//			customerService := service.INewCustomerService(&customerRepo, &userRepo, &tokenRepo, &tokenService, &mailService, &uuidService, &otpService)
//
//			// Actual method call
//			output, err := customerService.CreateUser(testCase.request)
//
//			if err != nil {
//				assert.Equal(t, testCase.expectedValueErrorResponse, err)
//			}
//			// Expected output
//			expected := testCase.expectedValueDataResponse
//
//			assert.Equal(t, expected, output)
//		})
//	}
//
//}

func TestCustomerService_ActivateUser(t *testing.T) {
	customerRepo := mocks.ICustomerRepository{}
	tokenService := mocks.ITokenService{}
	mailService := mocks.IMail{}
	uuidService := mocks.IUuid{}
	otpSerivce := mocks.IOtp{}

	testCases := []struct{
		name string

		tkRepoFindByUserTokenInput string
		tkRepoFindByUserTokenReturn *domain.Token
		tkRepoFindByUserTokenReturnErr error

		userRepoFindByIDInput int
		userRepoFindByIDReturn *domain.User
		userRepoFindByIDReturnErr error

		userRepoUpdateInput domain.User
		userRepoUpdateReturn *domain.User
		userRepoUpdateReturnErr error

		expectedReturn error
	}{
		{
			"Test activate user with valid token",
			"abcdefg",
			&domain.Token{UserId: 101, Token: "abcdefg", OTP: "123456", ExpiresOn: time.Now().Add(1 * time.Hour)},
			nil,
			101,
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: false},
			nil,
			domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			nil,
			nil,
		},
		{
			"Test activate user with invalid token",
			"",
			nil,
			errors.New("token does not exist"),
			101,
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: false},
			nil,
			domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			nil,
			errors.New("token does not exist"),
		},
		{
			"Test activate user with expired token",
			"abcdefg",
			&domain.Token{UserId: 102, Token: "abcdefg", OTP: "123456", ExpiresOn: time.Now().Add(-10 * time.Minute)},
			nil,
			102,
			&domain.User{Model: gorm.Model{ID: 102}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: false},
			nil,
			domain.User{Model: gorm.Model{ID: 102}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: false},
			&domain.User{Model: gorm.Model{ID: 102}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			nil,
			errors.New("token expired"),
		},
		{
			"Test activate user with valid token where user doesn't exist",
			"abcdefg",
			&domain.Token{UserId: 101, Token: "abcdefg", OTP: "123456", ExpiresOn: time.Now().Add(1 * time.Hour)},
			nil,
			101,
			nil,
			errors.New("user does not exist"),
			domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			nil,
			errors.New("user does not exist"),
		},
		{
			"Test activate user where user update fails",
			"abcdefg",
			&domain.Token{UserId: 101, Token: "abcdefg", OTP: "123456", ExpiresOn: time.Now().Add(1 * time.Hour)},
			nil,
			101,
			&domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: false},
			nil,
			domain.User{Model: gorm.Model{ID: 101}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way", IsActive: true},
			nil,
			errors.New("user update failed"),
			errors.New("user update failed"),
		},
	}

	for _, tC := range testCases{
		t.Run(tC.name, func(t *testing.T){
			tokenRepo := mocks.ITokenRepository{}
			tokenRepo.On("FindByToken", tC.tkRepoFindByUserTokenInput).Return(tC.tkRepoFindByUserTokenReturn, tC.tkRepoFindByUserTokenReturnErr)

			userRepo := mocks.IUserRepository{}
			userRepo.On("FindByID", tC.userRepoFindByIDInput).Return(tC.userRepoFindByIDReturn, tC.userRepoFindByIDReturnErr)
			userRepo.On("Update", tC.userRepoUpdateInput).Return(tC.userRepoUpdateReturn, tC.userRepoUpdateReturnErr)

			customerService := service.INewCustomerService(&customerRepo, &userRepo, &tokenRepo, &tokenService, &mailService, &uuidService, &otpSerivce)

			err := customerService.ActivateUser(tC.tkRepoFindByUserTokenInput)

			assert.Equal(t, tC.expectedReturn, err)
		})
	}
}