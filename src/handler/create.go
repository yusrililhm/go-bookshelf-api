package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func AddBook(c *gin.Context)  {
	var book model.Book
	if err := c.BindJSON(&book);err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": "gagal menambahkan buku : " + err.Error(),
		})
		return
	} else {
		model.Books = append(model.Books, book)
		c.JSON(200, gin.H{
			"status": "success",
			"message": "Buku berhasil ditambahkan",
			"book": map[string]any{
				"id": book.Id,
			},
		})
	}
}