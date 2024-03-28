package api

import (
	chatController "chat/controller"
	"github.com/gin-gonic/gin"
	videoController "video/controller"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/chat/sayhello", chatController.SayHello)

	//********	init video interfaces ********
	// video test
	engine.GET("/video/ping", videoController.Ping)
}
