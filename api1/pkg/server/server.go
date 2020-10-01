package server

import (
"ResnsBackend-api1/pkg/server/controller"
	"ResnsBackend-api1/pkg/middleware"
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
	Server.POST("auth/signin", middleware.HandleAuthenticate())
}