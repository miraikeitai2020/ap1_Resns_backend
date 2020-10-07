package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.ShouldBindJSON(&articleRequest)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		articles, err := model.GetArticles(articleRequest.Genre, articleRequest.Month, articleRequest.Year)
		fmt.Println(articles)
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, "")
	}
}

func HandleUserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK)
	}
}
