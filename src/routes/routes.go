package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/config"
	"github.com/yusrililhm/go-bookshelf-api/src/handler"
)

func Routes()  {
	r := gin.Default()

	// port

	port := config.GetEnv("PORT")

	// cors

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://search.opensuse.org/"},
		AllowMethods: []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders: []string{"Content-Type", "application/json"},
	}))

	book := r.Group("")
	{
		book.POST("/books", handler.AddBook)
		book.GET("/books", handler.ReadAllBook)
		book.GET("/books/:id", handler.ReadBookById)
		book.PUT("/books/:id", handler.UpdateBook)
		book.DELETE("/books/:id", handler.DeleteBook)
	}

	r.Run(":" + port)
}