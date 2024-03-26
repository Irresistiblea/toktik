package api

import (
	"chat/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/chat/sayhello", controller.SayHello)
}
