package service

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/jinzhu/gorm"
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

		mailServiceSendMailInput mail.EMailMessage
		mailServiceSendMailOutput string
		mailServiceSendMailError error

		expectedValueDataResponse *domain.Customer
		expectedValueErrorResponse error

	}{

	}


}