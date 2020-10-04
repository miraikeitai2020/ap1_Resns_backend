package controller

import (
	"ResnsBackend-api3/pkg/server/view"
	"ResnsBackend-api3/pkg/server/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var(
	request view.ArticleDetailRequest

)

func HandleAtricleDetailSend()gin.HandlerFunc {
	return func(c *gin.Context) {
		err:=c.ShouldBindJSON(&request)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		fmt.Println(request.ArticleID)
		articles, err := model.GetArticles(request.ArticleID);
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(articles.ArticleID)
		// // 生成した認証トークンを返却
		//c.JSON(http.StatusOK, view.Article{Article_id: article.Article_id, Image_path: article.Image_path,
		//	Title: article.Title, Context: article.Context,Genre:article.Genre, Nice: article.Nice,EraYear:article.EraYear,EraMonth:article.EraMonth})
		c.JSON(200,"hello ")
	}
}

func HandleAtricleSend()gin.HandlerFunc{
	return func(c *gin.Context) {
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}