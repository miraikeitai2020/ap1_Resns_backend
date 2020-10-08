package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ユーザー情報
func HandleUserSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := context.BindJSON(&HandleUserSet); err != nil {
			context.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		
		uuid, err := uuid.NewRandom()
		if err != nil {
			context.JSON(http.StatusInternalServerError, "AuthToken is error")
			return
		}
		restoken.Token = uuid.String()
		if err := HandleUserUpdate(HandleUserSet.Id, HandleUserSet.Name, HandleUserSet.Image, HandleUserSet.Year, HandleUserSet.Month HandleUserSet.Gender , restoken.Token); err != nil {
			context.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}	

}
//ユーザー情報更新
func HandleUserUpdate(id, name, image, token string, year, month, gender int) gin.HandlerFunc {
	return func(c *gin.Context) {
		stmt, err := db.Con.Prepare("INSERT INTO UserSet VALUES (?,?,?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(id, name, image, token, year, month, gender)
		return err
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}
