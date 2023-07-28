package main

import (
	"github.com/firmananshariadjie/tugas-akhir-api/controllers/eventcontroller"
	"github.com/firmananshariadjie/tugas-akhir-api/controllers/participantcontroller"
	"github.com/firmananshariadjie/tugas-akhir-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	
	r.GET("/api/events", eventcontroller.Index)	
	r.GET("/api/event/:id", eventcontroller.Show)
	r.POST("/api/event", eventcontroller.Create)
	r.PUT("/api/event/:id", eventcontroller.Update)
	r.DELETE("/api/event", eventcontroller.Delete)

	r.GET("/api/participants", participantcontroller.Index)
	r.GET("/api/participant/:id", participantcontroller.Show)
	r.POST("/api/participant", participantcontroller.Create)
	r.PUT("/api/participant/:id", participantcontroller.Update)
	r.DELETE("/api/participant", participantcontroller.Delete)

	r.Run()
}
