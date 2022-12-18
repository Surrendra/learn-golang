package main

import (
	"example/todo_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/product", ProductController.Index)
	r.GET("/api/product/:id", ProductController.Show)
	r.POST("/api/product", ProductController.Store)
	r.PUT("/api/product/:id", ProductController.Update)
	r.DELETE("/api/product/:id", ProductController.Delete)
	r.Run()
}
