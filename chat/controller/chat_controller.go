package controller

import (
	"chat/service"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	// get cid from context.
	cid := 123
	res := service.SayHelloService(cid)
	c.Writer.Write([]byte(res))
}
