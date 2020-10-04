package model

import (
	"ResnsBackend-api3/pkg/db"
	"database/sql"
	"log"
)

//  記事テーブルデータ
type Article struct{
	ArticleID string
	Title string
	ImagePath string
}
type ArticleDetails struct {
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


func GetArticles(genre,month,year string)([]Article,error){
	var articles []Article
	var article *Article
	var rows *sql.Rows
	var err error
	if(month==""){
		rows, err = db.Conn.Query("SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_year=?;",genre,year)
	}else if (year==""){
		rows, err = db.Conn.Query("SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_month=?;",genre,month)
	}else if (month==""&&year==""){
		rows, err = db.Conn.Query("SELECT article_id,title,image_path FROM articles_contents WHERE genre=?;",genre)
	}
	rows, err = db.Conn.Query("SELECT article_id,title,image_path FROM articles_contents WHERE genre=? AND era_year=? AND era_month=?;",genre,year,month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		article, err = convertToArticle(rows)
		if err != nil {
			return nil, err
		}
		articles=append(articles,*article)
	}
	return articles,nil
}

func convertToArticle(rows *sql.Rows) (*Article, error){
	article:=Article{}
	if err := rows.Scan(&article.ArticleID,&article.Title,&article.ImagePath); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}


func GetArticleDetail(articleID string)(*ArticleDetails,error){
	row := db.Conn.QueryRow("SELECT title,image_path,context,nice FROM articles_contents WHERE article_id=?",articleID)
	return convertToArticleDetails(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToArticleDetails(row *sql.Row) (*ArticleDetails, error) {
	var article ArticleDetails
	if err := row.Scan(&article.Title,&article.ImagePath,&article.Context,&article.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &article, nil
}