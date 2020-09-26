package server

import (
"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//アカウント作成
	Server.POST("auth/cerate", controller.HandleAuthCreate())
}