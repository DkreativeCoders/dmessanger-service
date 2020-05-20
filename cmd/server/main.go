package main

import (
	"fmt"
	_ "github.com/DkreativeCoders/dmessanger-service/doc"
	"github.com/DkreativeCoders/dmessanger-service/pkg"
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
