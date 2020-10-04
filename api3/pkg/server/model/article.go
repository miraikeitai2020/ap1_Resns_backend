package model

import (
	"ResnsBackend-api3/pkg/db"
	"database/sql"
	"log"
)

//  記事テーブルデータ
type Article struct {
	Title string
	ImagePath string
	Context string
	Nice string
}
//  コメントテーブルデータ
type Comment struct {
	UserName string
	UserImage string
	Contents string
}

var (
	articles Article
)


func GetArticles(articleID string)(*Article,error){
	row := db.Conn.QueryRow("SELECT title,image_path,context,nice FROM articles_contents WHERE article_id=?",articleID)
	return convertToArticle(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToArticle(row *sql.Row) (*Article, error) {
	if err := row.Scan(&articles.Title,&articles.ImagePath,&articles.Context,&articles.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	//fmt.Println(a)
	return &articles, nil
}