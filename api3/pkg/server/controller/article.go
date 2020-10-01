package controller

import (
	"ResnsBackend-api3/pkg/server/view"
	"ResnsBackend-api3/pkg/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
func HandleAtricleSend()gin.HandlerFunc {
	return func(c *gin.Context) {
		article, err := model.GetArticles();
		if err != nil {
			log.Fatal(err)
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.Article{Article_id: article.Article_id, Image_path: article.Image_path,
			Title: article.Title, Context: article.Context,Genre:article.Genre, Nice: article.Nice,EraYear:article.EraYear,EraMonth:article.EraMonth})
	}
}