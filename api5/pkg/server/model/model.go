package model

import (
	"ResnsBackend-api5/db"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type AddNice struct {
	Article_id string `json:"article_id"`
	User_id    string `json:"user_id"`
}
type Nices struct {
	Nice int `json:"nice"`
}

var (
	nice Nices
)

func GetAndUpdateNices(articleID string, userID string) (*Nices, error) {

	row := db.Conn.QueryRow("SELECT nice FROM articles_contents WHERE article_id=?", articleID)
	if err := row.Scan(&nice.Nice); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("AuthToken is error")
	}
	niceID := uuid.String()

	var checkNiceStatus string
	rows, err := db.Conn.Query("SELECT nice_id FROM articles_nice_status WHERE user_id=? AND article_id=?", userID, articleID)
	for rows.Next() {
		if err = rows.Scan(&checkNiceStatus); err != nil {
			log.Println(err)
		}
	}

	if checkNiceStatus == "" {
		stmts, err := db.Conn.Prepare("INSERT INTO articles_nice_status VALUES (?,?,?)")
		if err != nil {
			return nil, err
		}
		_, err = stmts.Exec(niceID, articleID, userID)

		nice.Nice = nice.Nice + 1
		stmts, err = db.Conn.Prepare("UPDATE articles_contents SET nice = ?  WHERE article_id=?")
		if err != nil {
			return nil, err
		}
		_, err = stmts.Exec(nice.Nice, articleID)

		row := db.Conn.QueryRow("SELECT nice FROM articles_contents WHERE article_id=?", articleID)
		if err := row.Scan(&nice.Nice); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			log.Println(err)
			return nil, err
		}
	}
	return &nice, nil

}
