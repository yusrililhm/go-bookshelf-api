package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server is running on port 9000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}