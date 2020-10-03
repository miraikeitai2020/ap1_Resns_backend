package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUserSet() gin.HandlerFunc {
	return func(c *gin.Context) {
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
