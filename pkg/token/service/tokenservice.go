package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"time"
)

func INewTokenService(tokenRepository irepository.ITokenRepository) iservice.ITokenService {
	return service{tokenRepository}
}

type service struct {
	tokenRepository irepository.ITokenRepository
}

func (s service) CreateTokenWithExpirationInHours(UserID uint, uniqueID string, expirationTimeInHour time.Duration) (*domain.Token, error) {
	token := domain.Token{}
	token.UserId = UserID
	token.Token = uniqueID
	token.ExpiresOn = time.Now().Add(expirationTimeInHour * time.Hour)

	newToken, err := s.tokenRepository.Create(token)

	if err != nil {
		return nil, errors.New("error occurred, try again")
	}
	return newToken, nil
}

func (s service) CreateTokenWithExpirationInMinutes(UserID uint, uniqueID string, expirationTimeInMinutes time.Duration) (*domain.Token, error) {
	panic("implement me")
}
