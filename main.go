package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"action": "pong",
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"action": "ping",
		})
	})
	r.Run()
}
