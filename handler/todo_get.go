package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-task/models"
	"test-task/routers"
)

func GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		db := routers.GetDB()
		items := []models.Item{}
		db.Find(&items)
		c.JSON(http.StatusOK, items)
	}
}

func GetById() func(c *gin.Context) {
	return func(c *gin.Context) {
		item := models.Item{}
		db := routers.GetDB()
		id := c.Param("id")
		db.First(&item, id)
		c.JSON(http.StatusOK, item)
	}
}
