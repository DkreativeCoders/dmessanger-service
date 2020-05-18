package main

import (
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg"
	_ "github.com/DkreativeCoders/dmessanger-service/doc"
	"log"
)


func main() {

	srv, db := pkg.NewServer()
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
