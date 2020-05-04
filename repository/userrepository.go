package repository

import (
	"github.com/danieloluwadare/dmessanger/interfaces"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/jinzhu/gorm"
)

func NewInMemoryRepository(db *gorm.DB) interfaces.IUserRepository {
	return userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetUser(id int) *models.User {
	var user models.User
	return &user
}
func (u userRepository) CreateUser(user models.User) error {
	return nil

}
