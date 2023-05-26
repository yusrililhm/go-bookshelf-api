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

	book := r.Group("book")
	{
		book.POST("/add", handler.AddBook)
		book.GET("/", handler.ReadAllBook)
		book.GET("/:id", handler.ReadBookById)
		book.PUT("/update{id}", handler.UpdateBook)
		book.DELETE("/delete{id}", handler.DeleteBook)
	}

	r.Run(":" + port)
}