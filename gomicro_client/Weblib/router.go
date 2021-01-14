package Weblib

import (
	"github.com/gin-gonic/gin"
	"gomicro_client/Services"
)

func NewGinRouter(prodService Services.ProdService)  *gin.Engine{

	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService),ErrorMiddeware())
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdList)
		v1Group.Handle("GET", "/prods/:pid", GetProdDetail)
	}

	return ginRouter

}