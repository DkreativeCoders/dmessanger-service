package orm

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/jinzhu/gorm"
)

//NewOrmUserRepository This returns an interface of the struct
func NewOrmUserRepository(db *gorm.DB) irepository.ICustomerRepository {
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

func (o ormCustomerRepository) Save(user domain.Customer) (*domain.Customer, error) {
	panic("implement me")
}

func (o ormCustomerRepository) Update(user domain.Customer) (*domain.Customer, error) {
	panic("implement me")
}


