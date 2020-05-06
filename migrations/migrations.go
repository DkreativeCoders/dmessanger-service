package migrations

import "github.com/danieloluwadare/dmessanger/models"

func init() {
	initiateModelMigration()
}

func initiateModelMigration() {
	// Migrate the schema
	GetDB().AutoMigrate(&models.User{})
}