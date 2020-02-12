package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleHandler(c *gin.Context) {
	// for test
	articles := []struct {
		Title    string
		Abstract string
		Category string
	}{{"article001", "some abs", "test"}, {"article002", "some abs", "test"},
		{"article003", "some abs", "test"}, {"article004", "some abs", "test"}}
	c.HTML(http.StatusOK, "user/article.html", gin.H{"articles": articles})
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.html", nil)
}
