package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"gomicro_client/Services"
	"strconv"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

func defaultProds(rsp interface{}) () {
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 1; i++ {
		models = append(models, newProd(i+2000, "prodname"+strconv.Itoa(2000+int(i))))
	}
	result := rsp.(*Services.ProdLIstResponse)
	result.Data=models
}

func defaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *Services.ProdDetailResponse:
		t.Data=newProd(10,"降级商品")
	case *Services.ProdLIstResponse:
		defaultProds(rsp)

	}
}

type ProdsWrapper struct {
	client.Client
}

func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error  {

	cmdName:=req.Service()+"."+req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout: 2000,
		RequestVolumeThreshold: 2,
		ErrorPercentThreshold: 50,
		SleepWindow: 5000,
	}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do("getprods", func() error {
		return this.Client.Call(ctx,req,rsp)
	}, func(e error) error {
		defaultData(rsp)
		return nil
	})

}

func NewLogWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}