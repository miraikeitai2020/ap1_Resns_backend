package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCommentSend()gin.HandlerFunc{
	return func(c *gin.Context) {
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}

func HandleCommentUpdate()gin.HandlerFunc{
	return func(c *gin.Context) {
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK,"" )
	}
}
