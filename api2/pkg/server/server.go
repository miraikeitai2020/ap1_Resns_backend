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
	//アカウントに基づくプロフィールの設定
	Server.POST("user/set",middleware.Authenticate(controller.HandleUserSet()))
	//アカウントに基づくプロフィールの情報更新
	Server.POST("user/update", middleware.Authenticate(controller.HandleUserUpdate()))
	//記事のデータ送信
	Server.GET("/article/send", controller.HandleAtricleSend())
	//記事のコメント送信
	Server.GET("/comment/send",controller.HandleCommentSend() )
	//記事のコメント更新
	Server.GET("/comment/update",controller.HandleCommentUpdate() )
	//記事のいいね数送信
	Server.GET("/nice/send", controller.HandleNiceSend())
	//記事のいいね数更新
	Server.GET("/nice/update", controller.HandleNiceUpdate())

}