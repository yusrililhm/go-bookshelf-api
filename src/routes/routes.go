package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/handler"
)

func Routes() *gin.Engine {
	r := gin.Default()

	// port

	book := r.Group("books")
	{
		book.POST("/", handler.AddBook)
		book.GET("/", handler.ReadAllBook)
		book.GET("/:id", handler.ReadBookById)
		book.PUT("/:id", handler.UpdateBook)
		book.DELETE("/:id", handler.DeleteBook)
	}

	return r
}