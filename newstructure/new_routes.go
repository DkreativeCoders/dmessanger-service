package newstructure

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/newstructure/migrations"
	"github.com/danieloluwadare/dmessanger/newstructure/user/controller/chttp"
	 "github.com/danieloluwadare/dmessanger/newstructure/user/repository/orm"
	"github.com/danieloluwadare/dmessanger/newstructure/user/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

func NewServer() (*http.Server, *gorm.DB) {

	//Get database connection
	dbConnection := migrations.GetDataBaseConnection()
	//Migrate all models
	migrations.InitiateModelMigration(dbConnection)
	//router created
	router := mux.NewRouter()




	//Initialize the repository for any the service
	userRepository := orm.NewOrmUserRepository(dbConnection)
	//Initialize the Service for any the handler
	userService := service.INewService(userRepository)
	//pass in the route and the user service
	chttp.NewUserHandler(router,userService)



	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	srv := &http.Server{Handler: router, Addr: ":"+port,}

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	return srv,dbConnection
}
