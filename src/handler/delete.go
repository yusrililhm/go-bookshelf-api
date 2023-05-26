package handler

import "github.com/gin-gonic/gin"

func DeleteBook(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": "success",
		"message": "Buku berhasil dihapus",
	})
}