package main

import (
	"fmt"
	"log"
)

func main() {

	srv, db := NewServer()
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
