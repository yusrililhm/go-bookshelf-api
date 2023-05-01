package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/routes"
)

func main() {
	routes.Routes()

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server is running on port 9000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}