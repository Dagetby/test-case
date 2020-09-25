package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-task/models"
	"test-task/routers"
)

func DeleteById() func(c *gin.Context) {
	return func(c *gin.Context) {
		item := models.Item{}
		db := routers.GetDB()
		id := c.Param("id")
		db.First(&item, id)

		if item.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "No todo found!",
			})
			return
		}

		db.Delete(&item)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Todo deleted successfully!",
		})
	}
}
