package main

import (
	"github.com/gin-gonic/gin"
	"test-task/handler"
)

func main() {

	r := gin.Default()

	r.GET("/todos", handler.GetAll())
	r.GET("/todos/:id", handler.GetById())
	r.POST("/todos", handler.CreateNew())
	r.PUT("/todos/:id", handler.UpdateById())
	r.DELETE("/todos/:id", handler.DeleteById())

	r.Run()

}
