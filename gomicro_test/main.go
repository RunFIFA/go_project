package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"myProject/gomicro_test/Services"
	"myProject/gomicro_test/ServicesImpl"
)

func main() {

	consulReg:=consul.NewRegistry(
		registry.Addrs("techad.top:8500"),
		)

	myService:=micro.NewService(
		micro.Name("api.luo.com.myapp"),
		micro.Address(":8001"),
		micro.Registry(consulReg),
		)

	Services.RegisterTestServiceHandler(myService.Server(), new(ServicesImpl.TestService))
	myService.Run()
}
