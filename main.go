package main

import (
	"github.com/amndnaziz/belajar-go-api/controller/productcontroller"
	"github.com/amndnaziz/belajar-go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.RUN()
}
