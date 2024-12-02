package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var app(
	app *gin.Engine
)

func routes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	routes(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}