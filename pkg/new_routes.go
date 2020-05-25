package pkg

import (
	"fmt"
	chttp2 "github.com/DkreativeCoders/dmessanger-service/pkg/customer/controller/chttp"
	customerOrm "github.com/DkreativeCoders/dmessanger-service/pkg/customer/repository/orm"
	customerService "github.com/DkreativeCoders/dmessanger-service/pkg/customer/service"
	"github.com/DkreativeCoders/dmessanger-service/pkg/migrations/gmorm"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/controller/chttp"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/repository/orm"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func NewServer() (*http.Server, *gorm.DB) {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dialect := os.Getenv("DB_CONNECTION")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//Get database connection
	dbConnection := gmorm.GetDataBaseConnection(dialect, username, password, dbName, dbHost, dbPort)
	gmorm.InitiateModelMigration(dbConnection)
	//Migrate all models
	gmorm.InitiateModelMigration(dbConnection)
	//router created
	router := mux.NewRouter()

	//Initialize the repository for any the service
	userRepository := orm.NewOrmUserRepository(dbConnection)
	//Initialize the Service for any the handler
	userService := service.INewService(userRepository)
	//pass in the route and the user service
	chttp.NewUserHandler(router, userService)


	//Initialize the repository for any the service
	customerRepository := customerOrm.NewOrmCustomerRepository(dbConnection)
	//Initialize the Service for any the handler
	customerService := customerService.INewCustomerService(customerRepository,userRepository)
	//pass in the route and the user service
	chttp2.NewCustomerHandler(router,customerService)




	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	srv := &http.Server{Handler: router, Addr: ":" + port}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	return srv, dbConnection
}
