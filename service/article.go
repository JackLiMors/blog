package service

import (
	"blog/dao/db"
	"blog/model"
	"fmt"
)

func GetArticleRecordList(pageNum, pageSize int) (
	articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

func GetArticleRecordListById(categoryId, pageNum, pageSize int) (
	articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Printf("get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}
	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	return
}

func InsertArticle(title, abstract, content string) (err error) {
	articleDetail := &model.ArticleDetail{}
	articleDetail.Title = title
	articleDetail.Summary = abstract
	articleDetail.Content = content
	// contentUtf8 := []rune(content)
	// minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	// articleDetail.Summary = string([]rune(content)[:minLength])
	id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert article succ, id:%d, err:%v\n", id, err)
	return
}
