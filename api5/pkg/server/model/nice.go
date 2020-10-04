package model

import (
	"ResnsBackend-api5/db"
	"database/sql"
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

func GetAndUpdateNices(article_id string) (*Nices, error) {
	row:= db.Conn.QueryRow("SELECT nice FROM articles_contents WHERE article_id=?", article_id)
	if err := row.Scan(&nice.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	nice.Nice=nice.Nice+1;
	stmt, err := db.Conn.Prepare("UPDATE articles_contents SET nice = ?")
	if err != nil {
		return nil,err
	}
	_, err = stmt.Exec(nice.Nice)
	return &nice, nil
}

//func insert(){
//	stmt, err :=db.Conn.Prepare(`INSERT INTO articles_nice_status VALUES(?,?,?)`)
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(niceID, articleID, userID)
//	return err
//}
