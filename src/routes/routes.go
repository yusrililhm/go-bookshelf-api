package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/handler"
)

func Routes()  {
	r := gin.Default()

	book := r.Group("book")
	{
		book.POST("/add", handler.AddBook)
		book.GET("/", handler.ReadAllBook)
		book.GET("/:id", handler.ReadBookById)
		book.PUT("/update{id}", handler.UpdateBook)
		book.DELETE("/delete{id}", handler.DeleteBook)
	}

	r.Run(":9000")
}