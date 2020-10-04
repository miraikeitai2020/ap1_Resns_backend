package controller

import (
	"ResnsBackend-api3/pkg/server/view"
	"ResnsBackend-api3/pkg/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var(
	request view.ArticleDetailRequest

)

func HandleAtricleSend()gin.HandlerFunc{
	return func(c *gin.Context) {
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}

func HandleAtricleDetailSend()gin.HandlerFunc {
	return func(c *gin.Context) {
		err:=c.ShouldBindJSON(&request)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		articles, err := model.GetArticles(request.ArticleID);
		if err != nil {
			log.Fatal(err)
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.Article{Title:articles.Title,ImagePath:articles.ImagePath,Context:articles.Context,Nice:articles.Nice})
		//c.JSON(200,"hello ")
	}
}
