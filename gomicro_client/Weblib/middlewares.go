package Weblib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomicro_client/Services"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys=make(map[string]interface{})
		context.Keys["prodservice"]=prodService
		context.Next()
	}
}

func ErrorMiddeware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r:=recover();r!=nil{
				context.JSON(500,gin.H{"status":fmt.Sprintf("%s",r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}