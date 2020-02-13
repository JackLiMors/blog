package db

import (
	"blog/model"
	"testing"
	"time"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "一般来说，是这样的。"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Summary = `第五魔法元素不存在哦`
	article.ArticleInfo.Title = "第五魔法元素在超光速场景下的边缘相对论效应"
	article.ArticleInfo.Username = "Mr.Sun"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}
	t.Logf("insert article succ, articleId:%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, len:%d\n", len(articleList))
}

func TestGetArticleInfo(t *testing.T) {
	articleInfo, err := GetArticleDetail(5)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, article:%#v\n", articleInfo)
}
