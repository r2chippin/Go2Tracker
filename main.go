package main

import (
	"Go2Tracker/controller"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// New a router
	router := gin.Default()

	// Load templates
	router.LoadHTMLGlob("resource/templates/*")

	// Handle main page
	controller.RouteMainPage(router)

	// Handel announce
	controller.RouteAnnounce(router)

	// Start server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// Server on
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Waiting signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	//Test only
}
