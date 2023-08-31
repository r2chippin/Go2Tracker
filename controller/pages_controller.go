package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Return index
func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go2Tracker",
	})
}

// Return a favicon if browser request
func handleFavicon(c *gin.Context) {
	c.File("resource/statics/favicon.ico")
}

// RouteMainPage  Handle main page
func RouteMainPage(r *gin.Engine) {
	r.GET("/", handleIndex)
	r.GET("/favicon.ico", handleFavicon)
}
