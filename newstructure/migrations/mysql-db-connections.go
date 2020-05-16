package migrations

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)


func GetDataBaseConnection() *gorm.DB{

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
	db := conn

	return db
}



// function to export global db object
//func GetDB() *gorm.DB {
//	return db
//}
