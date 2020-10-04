package model

import (
	"ResnsBackend-api3/pkg/db"
	"database/sql"
	"log"
)

//  記事テーブルデータ
type Article struct {
	ArticleID string
	ImagePath string
	Title string
	Context string
	Genre int
	Nice string
	EraYear int
	EraMonth int
}
//  コメントテーブルデータ
type Comment struct {
	UserName string
	UserImage string
	Cotents string
}

var (
	articles Article
)


func GetArticles(articleID string)(*Article,error){
	row := db.Conn.QueryRow("SELECT * FROM articles_contents WHERE article_id=?",2)
	return convertToArticle(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToArticle(row *sql.Row) (*Article, error) {
	article := Article{}
	if err := row.Scan(&article.ArticleID,&article.ImagePath,&article.Title, &article.Context,&article.Genre,&article.Nice,&article.EraYear,&article.EraMonth); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}