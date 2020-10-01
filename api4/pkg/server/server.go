package server

import (
	"ResnsBackend-api4/pkg/server/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//記事のコメント送信
	Server.GET("/comment/send",controller.HandleCommentSend() )
	//記事のコメント更新
	Server.GET("/comment/update",controller.HandleCommentUpdate() )
}