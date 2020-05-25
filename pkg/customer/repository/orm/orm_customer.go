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

//ormCustomerRepository and testing required
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

	//Create Transaction Object
	tx := o.db.Begin()
	//this functions runs at the end if there is a failure
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	//user := domain.User{
	//	FirstName: customer.FirstName,
	//	LastName: customer.LastName,
	//	Age: customer.Age,
	//	Email: customer.Email,
	//	PhoneNumber: customer.PhoneNumber,
	//	Password: customer.Password,
	//	Address: customer.Address,
	//}

	if dbc := tx.Create(&customer.User); dbc.Error != nil {
		tx.Rollback()
		return nil, dbc.Error
	}

	fmt.Println("user created =>", customer.User)
	fmt.Println("userID created =>", customer.User.ID)

	customer.UserId=customer.User.ID

	if dbc := tx.Create(&customer); dbc.Error != nil {
		tx.Rollback()
		return nil, dbc.Error
	}


	fmt.Println("user created =>", customer)

	if dbc := tx.Commit(); dbc.Error!=nil {
		tx.Rollback()
		return nil, dbc.Error

	}

	//return &customer
	return &customer, nil
}

func (o ormCustomerRepository) Update(customer domain.Customer) (*domain.Customer, error) {
	panic("implement me")
}


