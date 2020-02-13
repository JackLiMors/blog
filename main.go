package main

import (
	"blog/controller"
	"blog/dao/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	r.LoadHTMLGlob("html/**/*")
	r.GET("/index", controller.IndexHandler)
	r.Static("/static", "./static")
	r.GET("/article", controller.ArticleHandler)
	r.GET("/article/detail", controller.ArticleDetailHandler)
	r.GET("/writeArticle", controller.WriteArticleHandler)
	r.Run(":8000")
}
