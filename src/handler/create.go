package handler

import (
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func AddBook(c *gin.Context)  {
	var book model.Book

	id, _ := exec.Command("uuidgen").Output()
	t := time.Now()

	insertedAt := t.Format(time.RFC1123)
	updateAt := insertedAt

	book = model.Book{Id: string(id), InsertedAt: insertedAt, UpdateAt: updateAt}

	if err := c.BindJSON(&book);err != nil {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": "gagal menambahkan buku : " + err.Error(),
		})
		return
	} else {
		model.Books = append(model.Books, book)
		c.JSON(201, gin.H{
			"status": "success",
			"message": "Buku berhasil ditambahkan",
			"data": map[string]any{
				"bookId": book.Id,
			},
		})
	}
}