package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	fmt.Println(dbUrl)

	conn, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	initialMigration()
}

func initialMigration() {

	// Migrate the schema
	db.AutoMigrate(&User{})
}

// function to export global db object
func GetDB() *gorm.DB {
	return db
}
