package orm

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmTokenRepository(db *gorm.DB) irepository.ITokenRepository {
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

func (o ormTokenRepository) FindByToken(tk string) (*domain.Token, error) {
	var token domain.Token
	if dbc := o.db.Where("token = ?", tk).Find(&token); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &token, nil
}

func (o ormTokenRepository) UpdateToken(tk domain.Token) (*domain.Token, error) {
	var token domain.Token
	if dbc := o.db.Where("token = ?", tk.Token).Find(&token); dbc.Error != nil {
		return nil, dbc.Error
	}

	o.db.Save(tk)
	fmt.Println("Updated token")
	return &token, nil
}