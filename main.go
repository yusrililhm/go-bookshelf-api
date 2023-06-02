package main

import (
	"github.com/yusrililhm/go-bookshelf-api/src/config"
	"github.com/yusrililhm/go-bookshelf-api/src/routes"
)

func main() {

	// database setup

	db := config.ConnectDB()

	// server setup 

	port := config.GetEnv("PORT")
	r := routes.Routes(db)

	// running server
	r.Run(":" + port)
}