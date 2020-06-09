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
	mailService := mail.NewMailGunImplementationNoArgs()
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
	//corx := handlers.CORS(
	//	handlers.AllowedHeaders([]string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}),
	//	handlers.AllowedOrigins([]string{"*"}),
	//	handlers.AllowedMethods([]string{"POST, GET, OPTIONS, PUT, DELETE,OPTIONS"}),
	//	handlers.AllowCredentials(),
	//
	//)


	//(router)
	//c := cors.New(cors.Options{
	//	AllowedOrigins: []string{"*","http://foo.com", "http://foo.com:8080"},
	//	AllowCredentials: true,
	//	AllowedMethods: []string{"POST, GET, OPTIONS, PUT, DELETE,OPTIONS"},
	//	// Enable Debugging for testing, consider disabling in production
	//	Debug: true,
	//})

	//c.Handler(router)

	//router.Use(c)

	//router.Use(setUPR)
	//router.Use(loggingMiddleware)

	//router.Use(cors)

	//router.Use(mux.CORSMethodMiddleware(router))

	//srv := &http.Server{Handler:router, Addr: ":" + port}

	x:=handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)

	log.Fatal(http.ListenAndServe(":"+port, x))


	//log.Fatal(http.ListenAndServe(":"+port, (corx)(router)))


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

func setUPR(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		setupResponse(&w, r)

		next.ServeHTTP(w, r)
	})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://foo.com")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}