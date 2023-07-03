package main

import (
	"github.com/amndnaziz/belajar-go-api/controllers/productcontroller"
	"github.com/amndnaziz/belajar-go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/add-product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
