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
	articleDetailRequest view.ArticleDetailRequest
	articleRequest view.ArticleRequest
)

func HandleAtricleSend()gin.HandlerFunc{
	return func(c *gin.Context) {
		err:=c.ShouldBindJSON(&articleRequest)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		articles, err := model.GetArticles(articleRequest.Genre,articleRequest.Month,articleRequest.Year);
		fmt.Println(articles)
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}

func HandleAtricleDetailSend()gin.HandlerFunc {
	return func(c *gin.Context) {
		err:=c.ShouldBindJSON(&articleDetailRequest)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		article, err := model.GetArticleDetail(articleDetailRequest.ArticleID);
		if err != nil {
			log.Fatal(err)
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.Article{Title:article.Title,ImagePath:article.ImagePath,Context:article.Context,Nice:article.Nice})
		//c.JSON(200,"hello ")
	}
}
