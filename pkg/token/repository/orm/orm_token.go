package orm

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmUserRepository(db *gorm.DB) irepository.ITokenRepository {
	return ormTokenRepository{db}
}

type ormTokenRepository struct {
	db *gorm.DB
}

func (o ormTokenRepository) Create(token domain.Token) (*domain.Token, error) {
	// Create failed, do something e.g. return, panic etc.
	if dbc := o.db.Create(&token); dbc.Error != nil {
		return nil, dbc.Error
	}
	fmt.Println("token created =>", token)

	//return &token
	return &token, nil
}

func (o ormTokenRepository) FindByUserId(userId int) (*domain.Customer, error) {
	panic("implement me")
}
