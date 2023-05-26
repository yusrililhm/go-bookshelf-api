package routes

import "github.com/gin-gonic/gin"

func Routes()  {
	r := gin.Default()

	book := r.Group("book")
	{
		book.POST("/add", func(ctx *gin.Context) {})
		book.GET("/books", func(ctx *gin.Context) {})
		book.GET("/books{id}", func(ctx *gin.Context) {})
		book.PUT("/update{id}", func(ctx *gin.Context) {})
		book.DELETE("/delete{id}", func(ctx *gin.Context) {})
	}
}