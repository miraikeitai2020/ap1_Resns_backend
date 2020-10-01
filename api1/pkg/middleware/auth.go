package middleware

import (
"github.com/gin-gonic/gin"
	"net/http"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func HandleAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,"" )
	}
}
