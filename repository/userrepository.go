package repository

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/interfaces/irepository"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/jinzhu/gorm"
)

func NewUserRepository(db *gorm.DB) irepository.IUserRepository {
	return userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) FindByID(id int) *models.User {
	var user models.User
	return &user
}


func (u userRepository) FindAll() []models.User {
	users := make([]models.User, 0) // same as []int{0, 0}
	u.db.Find(&users)
	return users
}

//Save User or Return error
func (u userRepository) Save(user models.User) (*models.User, error) {
	// Create failed, do something e.g. return, panic etc.
	if dbc :=  u.db.Create(&user); dbc.Error != nil {
		return nil,dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", user)

	//return &user
	return &user,nil
}
