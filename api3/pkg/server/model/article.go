package model

import (
	"ResnsBackend-api3/pkg/db"
	"database/sql"
	"log"
)

//  articleテーブルデータ
type Article struct {
	Article_id string ` json:"article_id" `
	Image_path string `json:"image_path" `
	Title string `json:"title"`
	Context string `json:"context" `
	Genre    int `json:"genre" `
	Nice    int `json:"genre" `
	EraYear int  `json:"EraYear"`
	EraMonth int  `json:"EraMonth"`
}

func GetArticles()(*Article,error){
	row := db.Conn.QueryRow("SELECT * FROM articles WHERE article_id=?",2)
	return convertToArticle(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToArticle(row *sql.Row) (*Article, error) {
	article := Article{}
	if err := row.Scan(&article.Article_id, &article.Image_path,&article.Title, &article.Context,&article.Genre,&article.Nice,&article.EraYear,&article.EraMonth); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}