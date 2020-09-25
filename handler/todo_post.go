package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-task/models"
	"test-task/routers"
	"time"
)

func CreateNew() func(c *gin.Context) {
	return func(c *gin.Context) { // /post?id=1234&title="Hello World"&completed=false
		db := routers.GetDB()
		db.AutoMigrate(&models.Item{})

		completed, err := strconv.ParseBool(c.PostForm("completed"))
		if err != nil {
			fmt.Println("Error ON Convert String to bool", err)
		}
		item := models.Item{Title: c.PostForm("title"), Completed: completed, DueDate: time.Now()}
		db.Table("items").Save(&item)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "New notes created",
		})
	}
}
