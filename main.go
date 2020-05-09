package main

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/newstructure"
	"log"
)

func main() {

	srv, db := newstructure.NewServer()
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
