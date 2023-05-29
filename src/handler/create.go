package handler

import (
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/config"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func AddBook(c *gin.Context)  {
	var book model.Book

	// generate unique id

	id, _ := exec.Command("uuidgen").Output()
	t := time.Now()

	// date time format

	insertedAt := t.Format(time.RFC1123)
	updateAt := insertedAt

	book = model.Book{Id: string(id), InsertedAt: insertedAt, UpdateAt: updateAt}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(500, gin.H{
			"status": "fail",
			"message": "Buku gagal ditambahkan",
		})
		return
	} else if book.Name == "" {
		// response if name = ""
		
		c.JSON(400, gin.H{
			"status": "fail",
      "message": "Gagal menambahkan buku. Mohon isi nama buku",
		})
	} else if book.ReadPage > book.PageCount {
		// response if readPage > pageCount

		c.JSON(400, gin.H{
			"status": "fail",
			"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
	} else {
		// finished = readPage == pageCount

		finish := book.PageCount == book.ReadPage
		book.Finished = finish

		// response if add book success

		model.Books = append(model.Books, book)

		// insert into books

		db := config.ConnectDB()
		db.Create(&model.Books)

		c.JSON(201, gin.H{
			"status": "success",
			"message": "Buku berhasil ditambahkan",
			"data": map[string]any{
				"bookId": book.Id,
			},
		})
	}
}