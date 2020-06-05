package service

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/service"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomerService_CreateUser(t *testing.T) {

//	mock userRepository
//	mock customerRepository
//	mock tokenRepository
//	mock mailService

	testCases := []struct {
		name           string
		request  dto.CustomerRequest

		userRepoFindUserExistInput string
		userRepoFindUserExistReturnData bool

		customerRepositorySaveInput domain.Customer
		customerRepositorySaveReturnOutput *domain.Customer
		customerRepositorySaveReturnError error

		tokenRepositoryCreateInput domain.Token
		tokenRepositoryCreateReturnOutPut *domain.Token
		tokenRepositoryCreateReturnError error

		mailServiceSendMailInput *mail.EMailMessage
		mailServiceSendMailOutput string
		mailServiceSendMailError error

		expectedValueDataResponse *domain.Customer
		expectedValueErrorResponse error

	}{
		{
			"Test with valid customer request",
			dto.CustomerRequest{"Daniel","Dada","20","daniel@gmail.com","08282888", "password","address daniel"},
			"daniel@gmail.com",
			true,
			domain.Customer{User:domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},
			},
			&domain.Customer{User:domain.User{Model: gorm.Model{ID: 0},FirstName:"Daniel",LastName:"Dada",Age:"20",Email:"daniel@gmail.com", PhoneNumber:"08282888",Password:"password",Address:"address daniel"}},
			nil,
			domain.Token{UserId: 0,Token: "random string",ExpiresOn: time.Now().Add(1 * time.Hour)},
			&domain.Token{Model: gorm.Model{ID: 0},UserId: 0,Token: "random string",ExpiresOn: time.Now().Add(1 * time.Hour)},
			nil,

			mail.NewEMailMessage("DkreativeCoders Verify User","Please visit","daniel@gmail.com", nil),
			"success",
			nil,

			&domain.Customer{User:domain.User{Model: gorm.Model{ID: 0},FirstName:"Daniel",LastName:"Dada",Age:"20",Email:"daniel@gmail.com", PhoneNumber:"08282888",Password:"password",Address:"address daniel"}},
nil,
		},
	}


	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			userRepo := mocks.UserRepository{}
			userRepo.On("Save", testCase.repoInputData).Return(testCase.repoReturnData, testCase.repoReturnErr)


			// Create userService and inject mock repo
			userService := service.INewService(&userRepo)

			// Actual method call
			output, _ := userService.CreateUser(testCase.repoInputData)

			// Expected output
			expected := testCase.expectedVal

			assert.Equal(t, expected, output)
		})
	}


}