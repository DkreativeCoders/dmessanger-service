package migrations

import (
	"github.com/danieloluwadare/dmessanger/newstructure/domain"
	"github.com/jinzhu/gorm"
)


func InitiateModelMigration(db *gorm.DB)  *gorm.DB{
	// Migrate the schema
	db.AutoMigrate(&domain.User{})
	return  db
}