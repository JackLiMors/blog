package model

// `id``category_id``content``title``view_count``comment_count``username``status``summary``create_time``update_time`
import "time"

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Status       int       `db:"status"`
	Username     string    `db:"username"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
}

type ArticleRecord struct {
	ArticleInfo
	Category
}
