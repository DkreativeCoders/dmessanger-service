package migrations

import "github.com/danieloluwadare/dmessanger/oldstructure/models"

func init() {
	initiateModelMigration()
}

func initiateModelMigration() {
	// Migrate the schema
	GetDB().AutoMigrate(&models.User{})
}