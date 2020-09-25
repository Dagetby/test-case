package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-task/models"
	"test-task/routers"
)

func UpdateById() func(c *gin.Context) {
	return func(c *gin.Context) {
		item := models.Item{}
		db := routers.GetDB()
		id := c.Param("id")
		db.First(&item, id)

		if item.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
			return
		}

		c.JSON(http.StatusOK, item)
		db.Model(&item).Update("title", c.PostForm("title"))
		completed, _ := strconv.Atoi(c.PostForm("completed"))
		db.Model(&item).Update("completed", completed)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
	}
}
