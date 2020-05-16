// De-messenger
//
// De-messenger
//
//     Schemes: http
//     Host: localhost:8900
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
	"github.com/danieloluwadare/dmessanger/cmd"
	"log"
)

func main() {

	srv, db := cmd.NewServer()
	defer func() {
		fmt.Print("Closing Db")
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()


	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
