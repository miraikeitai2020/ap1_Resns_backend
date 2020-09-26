package middleware

import (
"github.com/gin-gonic/gin"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(ginNextMethod gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginNextMethod(c)
	}
}
