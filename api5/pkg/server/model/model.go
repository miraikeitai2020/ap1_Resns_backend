package model

import (
	"ResnsBackend-api5/db"
	"database/sql"
)

type AddNice struct {
	Article_id string `json:"article_id"`
	Token      string `json:"token"`
}
type Nices struct {
	Nice int `json:"nice"`
}

var (
	nice Nices
)

func GetNice(article_id string) (*Nices, error) {
	row, _ := db.Conn.Query("SELECT nice FROM articles_contents WHERE article_id=?", article_id)
	if err := row.Scan(&nice.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &nice, nil
}
