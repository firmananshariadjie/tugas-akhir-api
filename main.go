package main

import (
	"github.com/firmananshariadjie/tugas-akhir-api/controllers/eventcontroller"
	"github.com/firmananshariadjie/tugas-akhir-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// r.GET("/api/products", productcontroller.Index)
	// r.GET("/api/product/:id", productcontroller.Show)
	// r.POST("/api/product", productcontroller.Create)
	// r.PUT("/api/product/:id", productcontroller.Update)
	// r.DELETE("/api/product", productcontroller.Delete)
	
	r.GET("/api/events", eventcontroller.Index)
	r.GET("/api/event/:id", eventcontroller.Show)
	r.POST("/api/event", eventcontroller.Create)
	r.PUT("/api/event/:id", eventcontroller.Update)
	r.DELETE("/api/event", eventcontroller.Delete)

	r.Run()
}
