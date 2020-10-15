package model

import (
	"ResnsBackend-api2/db"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	Id     string `json:"id" `
	Name   string `json:"name" `
	Image  string `json:"image" `
	Year   int    `json:"year" `
	Month  int    `json:"month" `
	Gender int    `json:"gender" `
}

var (
	user *User
)

func Set(id, name, image string, year, month, gender int) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user(id, name, image, year, month, gender) VALUES(?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, name, image, year, month, gender)
	return err
}

func Update(id, name, image string, year, month, gender int) error {
	stmt, err := db.Conn.Prepare("UPDATE user SET name=?, image=?, year=?, month=?, gender=? WHERE id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, updErr := stmt.Exec(user.Name, user.Image, user.Year, user.Month, user.Gender, user.Id); updErr != nil {
		return parseError(updErr)
	}
	return nil
}

func parseError(err error) error {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("user not found")
		}
		return err
	}
	//以下,sql依存のエラー
	switch sqlErr.Number {
	case 1062:
		//email_UNIQUE error
		return fmt.Errorf("email already exist")
	}
	return err
}
