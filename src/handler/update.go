package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"gorm.io/gorm"
)

func UpdateBook(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var book model.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(404, gin.H{
			"status": "fail",
			"message": "Gagal memperbarui buku. Id tidak ditemukan",
		})
		return
	}

	var input model.Input

	if err := c.Bind(&input); err != nil {
		return
	} else if input.Name == "" {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": "Gagal memperbarui buku. Mohon isi nama buku",
		})
		return
	} else if input.ReadPage > input.PageCount {
		c.JSON(400, gin.H{
			"status": "fail",
			"message": "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	var m gorm.Model

	m.CreatedAt = time.Now()

	book.UpdateAt = m.UpdatedAt
	db.Model(&book).Updates(input)
	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil diperbarui",
	})
}