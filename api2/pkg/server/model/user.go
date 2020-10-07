package model

import (
	"ResnsBackend-api2/pkg/db"
	"database/sql"
	"log"
)

//  userテーブルデータ
type User struct {
	Id string `json:"id" `
	Name string `json:"name" `
	Image string `json:"image" `
	Year int `json:"year" `
	month int `json:"month" `
	Gender int `json:"gender" `
}

type Retoken struct {
	Token string `json:"token" `
}

func (ctrl *Controller) SetUser(context *gin.Context) {
	if err := context.BindJSON(&SetUser); err != nil {
		context.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	
	uuid, err := uuid.NewRandom()
	if err != nil {
		context.JSON(http.StatusInternalServerError, "AuthToken is error")
		return
	}
	restoken.Token = uuid.String()
	if err := UpdateUser(SetUser.Id, SetUser.Name, SetUser.Image, SetUser.Year, SetUser.Month SetUser.Gender , restoken.Token); err != nil {
		context.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	context.JSON(200, restoken)
}

func UpdateUser(id, name, image, token string, year, month, gender int) error {
	stmt, err := db.Con.Prepare("INSERT INTO signUp VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, name, token, year, month, gender)
	return err