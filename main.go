package main

import "log"

func main() {

	srv := NewServer()
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
