package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/clicked", func(c *gin.Context) {

		c.HTML(http.StatusOK, "simple_text.html", gin.H{
			"title": "Main website",
		})

	})

	r.POST("/hover", func(c *gin.Context) {
		c.HTML(http.StatusOK, "alert.html", gin.H{})
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		c.HTML(http.StatusOK, "search.html", gin.H{
			"item": query,
		})
	})

	r.Run()
}
