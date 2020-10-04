package controller

import (
	"ResnsBackend-api5/pkg/server/model"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	addNice model.AddNice
	nice    model.Nices
)

func HandleNiceSend() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&addNice); err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		nice, err := model.GetAndUpdateNices(addNice.Article_id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		if nice == nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"certification": "false"})
		}
		// 生成した認証トークンを返却
		c.JSON(http.StatusOK, nice)
	}
}

func HandleNiceUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&addNice); err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, "")
	}
}
