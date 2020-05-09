package orm

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/newstructure/domain"
	"github.com/danieloluwadare/dmessanger/newstructure/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmUserRepository(db *gorm.DB) irepository.IUserRepository {
	return ormUserRepository{db}
}

type ormUserRepository struct {
	db *gorm.DB
}


func (u ormUserRepository) FindByID(id int) *domain.User {
	var user domain.User
	return &user
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
	if dbc :=  u.db.Create(&user); dbc.Error != nil {
		return nil,dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", user)

	//return &user
	return &user,nil
}
