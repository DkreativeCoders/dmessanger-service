package migrations

import (
	"github.com/danieloluwadare/dmessanger/newstructure/domain"
	"github.com/jinzhu/gorm"
)


func InitiateModelMigration(db *gorm.DB)  {
	// Migrate the schema
	db.AutoMigrate(
		&domain.User{},
		&domain.Courier{},
		&domain.Customer{},
		)
//	define foreign key relationship here
	db.Model(&domain.Courier{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&domain.Customer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

}