package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/danieloluwadare/dmessanger/interfaces"

)

func NewInMemoryRepository(db *gorm.DB) interfaces.IUserRepository {
	return userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetUser(id int) (*models.User, error) {
	return nil nil
}
func (u userRepository) CreateUser(user models.User) error {
	return nil

}
