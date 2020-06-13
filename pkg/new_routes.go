package pkg

import (
	"fmt"
	mail "github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/uuid"
	chttp2 "github.com/DkreativeCoders/dmessanger-service/pkg/customer/controller/chttp"
	customerOrm "github.com/DkreativeCoders/dmessanger-service/pkg/customer/repository/orm"
	customerService "github.com/DkreativeCoders/dmessanger-service/pkg/customer/service"
	"github.com/DkreativeCoders/dmessanger-service/pkg/migrations/gmorm"
	tokenOrm "github.com/DkreativeCoders/dmessanger-service/pkg/token/repository/orm"
	token_Service "github.com/DkreativeCoders/dmessanger-service/pkg/token/service"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/controller/chttp"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/repository/orm"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/service"
	"github.com/gorilla/handlers"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
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
	//Migrate all models
	gmorm.InitiateModelMigration(dbConnection)
	//router created
	router := mux.NewRouter()

	//initialize configs
	apiKey := os.Getenv("SENDGRID_API_KEY")
	mailService := mail.NewSendGrid(apiKey)
	uuid := uuid.INewUuid()

	//Initialize Token Repository

	tokenRepository := tokenOrm.NewOrmTokenRepository(dbConnection)
	tokenService := token_Service.INewTokenService(tokenRepository)

	//Initialize the repository for any the service
	userRepository := orm.NewOrmUserRepository(dbConnection)
	//Initialize the Service for any the handler
	userService := service.INewService(userRepository)
	//pass in the route and the user service
	chttp.NewUserHandler(router, userService)

	//Initialize the repository for any the service
	customerRepository := customerOrm.NewOrmCustomerRepository(dbConnection)
	//Initialize the Service for any the handler
	newCustomerService := customerService.INewCustomerService(customerRepository, userRepository, tokenRepository, tokenService, mailService, uuid)
	//pass in the route and the user service
	chttp2.NewCustomerHandler(router, newCustomerService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	//router.Use(loggingMiddleware)


	corx:=handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)

	log.Fatal(http.ListenAndServe(":"+port, corx))


	//err := http.ListenAndServe(":"+port, (corx)(router)) //Launch the app, visit localhost:8000/api
	//if err != nil {
	//	fmt.Print(err)
	//}

	//return srv, dbConnection
	return nil, dbConnection

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			log.Println("okay options")

			return
		}
		next.ServeHTTP(w, r)
	})
}

