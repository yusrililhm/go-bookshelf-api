package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yusrililhm/go-bookshelf-api/src/config"
	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func DeleteBook(c *gin.Context)  {
	id := c.Param("id")
	for _, v := range model.Books {
		if v.Id == id{
			db := config.ConnectDB()
			db.Where("id = ?", v.Id).Delete(&v)
			c.JSON(200, gin.H{
				"status": "success",
				"message": "Buku berhasil dihapus",
			})
			c.Redirect(200, "/books/:id")
			return
		}
	}
	c.JSON(404, gin.H{
		"status": "fail",
		"message": "Buku gagal dihapus. Id tidak ditemukan",
	})
}