package main

import "log"

func main() {

	srv := Server()
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
