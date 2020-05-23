package orm

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmCustomerRepository(db *gorm.DB) irepository.ICustomerRepository {
	return ormCustomerRepository{db}
}

type ormCustomerRepository struct {
	db *gorm.DB
}

func (o ormCustomerRepository) FindByID(id int) (*domain.Customer, error) {
	panic("implement me")
}

func (o ormCustomerRepository) FindByUserId(userId int) (*domain.Customer, error) {
	panic("implement me")
}

func (o ormCustomerRepository) FindByEmail(email string) (*domain.Customer, error) {
	panic("implement me")
}

func (o ormCustomerRepository) FindAll() []domain.Customer {
	panic("implement me")
}

func (o ormCustomerRepository) Save(customer domain.Customer) (*domain.Customer, error) {
	if dbc := o.db.Create(&customer); dbc.Error != nil {
		return nil, dbc.Error
	}
	//u.db.Where("email = ?", user.Email).First(&newUser)
	fmt.Println("user created =>", customer)

	//return &user
	return &customer, nil
}

func (o ormCustomerRepository) Update(customer domain.Customer) (*domain.Customer, error) {
	panic("implement me")
}


