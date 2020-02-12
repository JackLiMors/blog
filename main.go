package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/**/*")
	r.Static("/static", "./static")
	r.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/article.html", gin.H{})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{})
	})
	r.Run(":8000")
}
