package db

import (
	"blog/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlstr := `insert into 
					article(content, summary, title, username, 
						category_id, view_count, comment_count)
				values(?, ?, ?, ?, ?, ?, ?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleId)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
							id, summary, title, view_count, content,
							 create_time, comment_count, username, category_id
						from 
							article 
						where 
							id = ?
						and
							status = 1
						`
	err = DB.Get(articleDetail, sqlstr, articleId)
	return
}

func GetArticleListByCategoryId(categoryId, pageNum,
	pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
						and
						category_id = ?
					order by create_time desc
					limit ?, ?`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}


func IsArticleExist(articleId int64) (exists bool, err error) {
	var id int64
	sqlstr := "select id from article where id=?"
	err = DB.Get(&id, sqlstr, articleId)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}
