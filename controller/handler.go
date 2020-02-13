package controller

import (
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ArticleHandler(c *gin.Context) {
	// for test
	ArticleRecordList, err := service.GetArticleRecordList(0, 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "special/500.html", nil)
	}
	c.HTML(http.StatusOK, "user/article.html", gin.H{"articles": ArticleRecordList})
}

func ArticleDetailHandler(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "special/500.html", nil)
		return
	}
	var articleInformation = make(map[string]interface{}, 10)
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "special/500.html", nil)
	}
	articleInformation["articleDetail"] = articleDetail
	c.HTML(http.StatusOK, "user/articleDetail.html", articleInformation)

}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.html", nil)
}

func WriteArticleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/writeArticle.html", nil)
}
