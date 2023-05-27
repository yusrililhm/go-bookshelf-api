package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func ReadAllBook(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": "success",
		"data": map[string]any{
			"books": model.Books,
		},
	})
}

func ReadBookById(c *gin.Context)  {
	id := c.Param("id")
	for _, v := range model.Books {
		if v.Id == id {
			c.JSON(200, gin.H{
				"status": "success",
				"data": map[string]any{
					"book": v,
				},
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"status": "fail",
		"message": "Buku tidak ditemukan",
	})
}