package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "index",
		})
	})

	server := web.NewService(
		web.Address("127.0.0.1:8001"),
		web.Handler(ginRouter),
	)

	server.Run()

}
