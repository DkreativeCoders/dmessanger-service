package repository

import (
	"fmt"
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

//Save And retrieves Saved user
func (u userRepository) Save(user models.User) (*models.User, error) {
	if dbc :=  u.db.Create(&user); dbc.Error != nil {
		// Create failed, do something e.g. return, panic etc.
		return nil,dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", user)

	//return &user
	return &user,nil
}
