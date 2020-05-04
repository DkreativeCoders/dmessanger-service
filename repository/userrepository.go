package repository

import (
	"github.com/danieloluwadare/dmessanger/interfaces/irepository"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/jinzhu/gorm"
)

func NewInMemoryRepository(db *gorm.DB) irepository.IUserRepository {
	return userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) FindByID(id int) *models.User {
	var user models.User
	return &user
}
func (u userRepository) Save(user models.User) error {
	return nil

}
