package handler

import "github.com/gin-gonic/gin"

func ReadAllBook(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil ditambahkan",
	})
}

func ReadBookById(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil ditambahkan",
	})
}