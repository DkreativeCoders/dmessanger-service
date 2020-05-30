package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/service"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateUser(t *testing.T) {

	testCases := []struct {
		name           string
		repoInputData  domain.User
		repoReturnData *domain.User
		repoReturnErr  error
		expectedVal    map[string]interface{}
	}{
		{
			"Test with valid user input",
			domain.User{FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			&domain.User{Model: gorm.Model{ID: 0}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			nil,
			map[string]interface{}{
				"status":  true,
				"message": "success",
				"data":    &domain.User{Model: gorm.Model{ID: 0}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			},
		},
		{
			"Test with error from repository",
			domain.User{Model: gorm.Model{ID: 0}, FirstName: "Adam", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			nil,
			errors.New("user already exist"),
			map[string]interface{}{
				"status":        false,
				"message":       "Error",
				"error_message": errors.New("user already exist"),
			},
		},
		{
			"Test with empty first name",
			domain.User{FirstName: "", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			&domain.User{Model: gorm.Model{ID: 0}, FirstName: "", LastName: "Mark", Age: "24", Email: "amark@gmail.com", PhoneNumber: "01-2345-6789", Password: "password", Address: "401, Hebert Mark Way"},
			nil,
			map[string]interface{}{
				"status":  false,
				"message": "User First name should be on the payload",
			},
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
