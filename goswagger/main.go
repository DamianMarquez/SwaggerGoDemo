package main

import "log"

func main() {
	log.Println("Hola")
	server := NewServer()
	server.ListenAndServe(":8080")
}
