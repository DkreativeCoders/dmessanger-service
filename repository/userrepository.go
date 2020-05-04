package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/danieloluwadare/dmessanger/interface"

)

func NewInMemoryRepository(db *gorm.DB) interface.IUserRepository {
	return userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetUser(id int) (*models.User, error){

}
func (u userRepository) CreateUser(user models.User) error {
	
}