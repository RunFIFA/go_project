package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro_client/Services"
	"gomicro_client/Weblib"
	"gomicro_client/Wrappers"
)


type logWrapper struct {
	client.Client
}

func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error  {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx,req,rsp)

}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("techad.top:8500"))

	myService:=micro.NewService(
		micro.Name("prodservive.client"),
		micro.WrapClient(NewLogWrapper),
		micro.WrapClient(Wrappers.NewLogWrapper),
	)

	prodService:=Services.NewProdService("prodservice",myService.Client())

	httpServier := web.NewService(
		web.Name("httpprodservice"),
		web.Address("127.0.0.1:8001"),
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	httpServier.Init()
	httpServier.Run()
}
