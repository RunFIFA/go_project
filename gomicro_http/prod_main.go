package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/Helper"
	"gomicro/ProdService"
	"log"
)

func main() {

	ginRouter := gin.Default()

	consulReg := consul.NewRegistry(
		registry.Addrs("techad.top:8500"),
	)

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			var pr Helper.ProdsRequest
			err:=context.Bind(&pr)
			if err != nil || pr.Size <= 0{
				pr = Helper.ProdsRequest{Size: 2}
			}
			context.JSON(200,
				gin.H{
					"data":ProdService.NewProdList(pr.Size)})
		})
	}

	server := web.NewService(
		web.Name("prodservise"),
		web.Address("127.0.0.1:8003"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	err := server.Init()
	if err != nil{
		log.Fatal(err)
	}

	server.Run()

}
