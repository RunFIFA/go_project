package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro_grpc/ServiceImpl"
	"gomicro_grpc/Services"
)

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("techad.top:8500"),
	)

	prodService:=micro.NewService(
		micro.Name("prodservice"),
		micro.Address("127.0.0.1:8011"),
		micro.Registry(consulReg),
	)

	prodService.Init()
	Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))
	prodService.Run()
}