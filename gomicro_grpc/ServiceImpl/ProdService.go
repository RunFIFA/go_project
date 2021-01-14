package ServiceImpl

import (
	"context"
	"fmt"
	"gomicro_grpc/Services"
	"strconv"
)

type ProdService struct {

}

func newProd(id int32, pname string) *Services.ProdModel{
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

func (*ProdService)  GetProdLIst(ctx context.Context, in *Services.ProdsRequest, out *Services.ProdLIstResponse) error {
	//time.Sleep(time.Second*3)
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0; i < in.Size; i++{
		models = append(models, newProd( i+1000 ,"prodname"+strconv.Itoa(1000+int(i))) )
	}
	fmt.Println(models)
	out.Data = models
	return nil
}

func (*ProdService) GetProdDetail(ctx context.Context, in *Services.ProdsRequest, out *Services.ProdDetailResponse) error {
	//time.Sleep(time.Second*3)
	out.Data=newProd(in.ProdId,"测试商品")
	return nil
}