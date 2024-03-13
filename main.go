package main

import (
	"github.com/Rezapahlevi3108/hacktiv-assignment-2/controllers/productcontroller"
	"github.com/Rezapahlevi3108/hacktiv-assignment-2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:item_id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:item_id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
