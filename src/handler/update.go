package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func UpdateBook(c *gin.Context)  {
	id := c.Param("id")
	for _, v := range model.Books {
		if v.Id == id {
			c.JSON(200, gin.H{
				"status": "success",
				"message": "Buku berhasil diupdate",
			})
		}
	}
}