package main

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/**/*")
	r.Static("/static", "./static")
	r.GET("/article", controller.ArticleHandler)
	r.GET("/index", controller.IndexHandler)
	r.Run(":8000")
}
