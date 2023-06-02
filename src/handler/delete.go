package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"gorm.io/gorm"
)

func DeleteBook(c *gin.Context)  {

	var books model.Book

	id := c.Param("id")

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).First(&books).Error; err != nil {
		c.JSON(404, gin.H{
			"status": "fail",
			"message": "Buku gagal dihapus. Id tidak ditemukan",
		})
		return
	}

	db.Delete(&books, "id = ?", id)

	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil dihapus",
	})

}