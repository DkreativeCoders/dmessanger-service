package gmorm

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/pkg/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func GetDataBaseConnection(username string, password string, dbName string, dbHost string) *gorm.DB{

	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	fmt.Println(dbUrl)

	conn, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Print(err)
	}
	db := conn

	return db
}

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



// function to export global db object
//func GetDB() *gorm.DB {
//	return db
//}
