package orm

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmUserRepository(db *gorm.DB) irepository.IUserRepository {
	return ormUserRepository{db}
}

type ormTokenRepository struct {
	db *gorm.DB
}

func (u ormUserRepository) FindByID(id int) (*domain.User, error) {

	var user domain.User
	dbc := u.db.Where(domain.User{Model: gorm.Model{
		ID: uint(id),
	}}).First(&user)

	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &user, nil
}

//FindAll Users
func (u ormUserRepository) FindAll() []domain.User {
	users := make([]domain.User, 0) // same as []int{0, 0}
	u.db.Find(&users)
	return users
}

//Save User or Return error
