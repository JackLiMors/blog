package controller

import (
	"blog/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	// input := []byte(articleDetail.Content)
	// output := blackfriday.Run(input)
	// html := bluemonday.UGCPolicy().SanitizeBytes(output)
	// articleDetail.Content = string(html)
	articleInformation["articleDetail"] = articleDetail
	c.HTML(http.StatusOK, "user/articleDetail.html", articleInformation)
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.html", nil)
}

func WriteArticleHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/writeArticle.html", nil)
}

func UploadArticleHandler(c *gin.Context) {
	c.Request.ParseForm()
	title := c.PostForm("title")
	abstract := c.PostForm("abstract")
	content := c.PostForm("content")
	// filename := "static/mddoc/"+title+".md"
	// file,err := os.Open(filename)
	// if err != nil {
	// 	fmt.Println("open file err=", err)
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(file)
	// writer.WriteString(content)
	// writer.Flush()
	service.InsertArticle(title, abstract, content)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/article")
}
