package server

import (
	"ResnsBackend-api3/pkg/server/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//記事のデータ送信
	Server.GET("/article/send", controller.HandleAtricleSend())
	Server.GET("/article/detailSend", controller.HandleAtricleDetailSend())
}