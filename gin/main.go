package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){

	engine := gin.Default()

	engine.Handle("POST", "/login", func(context *gin.Context) {

		fmt.Println(context.FullPath())
		//userName
		username, _ := context.GetQuery("as")
		fmt.Println(username)

		//passWord
		password, _ := context.GetQuery("pwd")
		fmt.Println(password)

		context.Writer.Write([]byte("User login"+username))
	})

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		//获取请求接口
		fmt.Println(context.FullPath())
		//获取字符串参数
		name := context.DefaultQuery("name", "")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello ," + name))
	})



	engine.Run()
}
