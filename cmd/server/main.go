// De-messenger
//
// De-messenger
//
//     Schemes: http,https,127.0.0.1
//     Host: dmessanger-service.herokuapp.com
//     Version: 0.0.1
//	   BasePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg"
	"log"
)

func main() {

	_, db := pkg.NewServer()
	defer func() {
		fmt.Print("Closing Db")
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//log.Println("Server listening on", srv.Addr)
	//log.Fatal(srv.ListenAndServe())
}
