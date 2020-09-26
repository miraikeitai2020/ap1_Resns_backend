package server

import (
	"ResnsBackend-api5/pkg/server/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//記事のいいね数送信
	Server.GET("/nice/send", controller.HandleNiceSend())
	//記事のいいね数更新
	Server.GET("/nice/update", controller.HandleNiceUpdate())
}