package server

import (
	"ResnsBackend-api6/pkg/server/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//お気に入りリストの記事追加
	Server.POST("/list/add", controller.HandleListAdd())
	//お気に入りリストの記事削除
	Server.POST("/list/delete", controller.HandleListDelete())
}