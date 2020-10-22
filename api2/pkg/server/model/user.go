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
//user *User
)

func Set(user User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO users(id,name,image,year,month,gender) VALUES(?,?,?,?,?,?)")
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("5555")
	_, err = stmt.Exec(user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender)
	return err
}

func Update(user User) error {
	stmt, err := db.Conn.Prepare("UPDATE users SET name=?, image=?, year=?, month=?, gender=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, updErr := stmt.Exec(user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender); updErr != nil {
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
		return fmt.Errorf("")
	}
	return err
}
