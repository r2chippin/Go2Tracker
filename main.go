package main

import (
	"Go2Tracker/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// New a router
	router := gin.Default()

	// Load templates
	router.LoadHTMLGlob("resource/templates/*")

	// Handle main page
	controller.HandleMainPage(router)

	// Start router
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
}
