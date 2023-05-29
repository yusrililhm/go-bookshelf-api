package main

import (
	"github.com/yusrililhm/go-bookshelf-api/src/config"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"github.com/yusrililhm/go-bookshelf-api/src/routes"
)

func main() {
	// server setup 

	port := config.GetEnv("PORT")
	r := routes.Routes()

	// database migration

	db := config.ConnectDB()
	db.AutoMigrate(&model.Book{})

	// running server
	r.Run(":" + port)
}