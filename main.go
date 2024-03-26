package main

import (
	"api"
	"fmt"

	"github.com/gin-gonic/gin"
)

func myMiddleWare(c *gin.Context) {
	fmt.Println("myMiddleWare")
}

func main() {
	engine := gin.Default()
	engine.Use(myMiddleWare)

	api.InitRouter(engine)

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
