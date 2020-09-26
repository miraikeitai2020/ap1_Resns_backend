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
	//アカウントに基づくプロフィールの設定
	Server.POST("user/set",middleware.Authenticate(controller.HandleUserSet()))
	//アカウントに基づくプロフィールの情報更新
	Server.POST("user/update", middleware.Authenticate(controller.HandleUserUpdate()))
}