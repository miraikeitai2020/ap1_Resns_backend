package controller

import (
	"ResnsBackend-api2/pkg/server/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	user model.User
)

//ユーザー情報
func HandleUserSet() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.ShouldBindJSON(&user); err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			context.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		uuid, err := uuid.NewRandom()
		if err != nil {
			context.JSON(http.StatusInternalServerError, "AuthToken is error")
			return
		}
		user.Id = uuid.String()

		if err := model.Set(user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender); err != nil {
			context.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// // 生成した認証トークンを返却
		context.JSON(http.StatusOK, model.User{user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender})
	}
}

//ユーザー情報更新
func HandleUserUpdate() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.BindJSON(&user); err != nil {
			context.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		if err := model.Update(user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender); err != nil {
			context.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// // 生成した認証トークンを返却
		context.JSON(http.StatusOK, model.User{user.Id, user.Name, user.Image, user.Year, user.Month, user.Gender})
	}
}
