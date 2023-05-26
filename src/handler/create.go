package handler

import "github.com/gin-gonic/gin"

func AddBook(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil ditambahkan",
	})
}