package controller

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// get cid from context.
	_, err := c.Writer.Write([]byte("Pong"))
	if err != nil {
		return
	}
}
