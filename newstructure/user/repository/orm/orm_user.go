package repository

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/newstructure/domain/irepository"
	"github.com/danieloluwadare/dmessanger/oldstructure/models"
	"github.com/jinzhu/gorm"
)

//INewUserRepository This returns an interface of the struct
func INewUserRepository(db *gorm.DB) irepository.IUserRepository {
	return ormUserRepository{db}
}

type ormUserRepository struct {
	db *gorm.DB
}


func (u ormUserRepository) FindByID(id int) *models.User {
	var user models.User
	return &user
}

//FindAll Users
func (u ormUserRepository) FindAll() []models.User {
	users := make([]models.User, 0) // same as []int{0, 0}
	u.db.Find(&users)
	return users
}

//Save User or Return error
func (u ormUserRepository) Save(user models.User) (*models.User, error) {
	// Create failed, do something e.g. return, panic etc.
	if dbc :=  u.db.Create(&user); dbc.Error != nil {
		return nil,dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", user)

	//return &user
	return &user,nil
}
