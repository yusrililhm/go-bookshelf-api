package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"gorm.io/gorm"
)

func ReadAllBook(c *gin.Context)  {

	books := []struct{
		Id string `json:"id"`
		Name string `json:"name"`
		Publisher string `json:"publisher"`
	}{}

	db := c.MustGet("db").(*gorm.DB)
	db.Model(&model.Book{}).Find(&books)
	c.JSON(200, gin.H{
		"status": "success",
		"data": map[string]any{
			"books": books,
		},
	})
}

func ReadBookById(c *gin.Context)  {

	var books model.Book

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.First(&books, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{
			"status": "fail",
			"message": "Buku tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data": map[string]any{
			"book": books,
		},
	})
}