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

type ormUserRepository struct {
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
func (u ormUserRepository) Save(user domain.User) (*domain.User, error) {
	// Create failed, do something e.g. return, panic etc.
	if dbc := u.db.Create(&user); dbc.Error != nil {
		return nil, dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", user)

	//return &user
	return &user, nil
}

//Update User or Return error
func (u ormUserRepository) Update(user domain.User) (*domain.User, error) {
	// Update failed, do something e.g. return, panic etc.
	// userDomain act as table name
	var userDomain domain.User

	dbc := u.db.Where(domain.User{Model: gorm.Model{
		ID: uint(user.ID),
	}}).First(&userDomain)

	if dbc.Error != nil {
		return nil, dbc.Error
	}

	u.db.Save(user)
	fmt.Println("Updated user")
	//return &user
	return &user, nil
}

func (u ormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	u.db.Where(&domain.User{Email: email}).First(&user)
	return &user, nil
}

func (u ormUserRepository) FindUserExist(email string) bool {
	var count int
	u.db.Model(&domain.User{}).Where("email = ?", email).Count(&count)
	fmt.Println("count user =>", count)
	if count <= 0 {
		return false
	} else {
		return true
	}

}
