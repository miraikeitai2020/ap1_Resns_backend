package model

import (
	"ResnsBackend-api5/db"
	"database/sql"
	"fmt"
	"log"
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
	row:= db.Conn.QueryRow("SELECT nice FROM articles_contents WHERE article_id=?", article_id)
	if err := row.Scan(&nice.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &nice, nil
}
