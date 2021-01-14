package Weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	"gomicro_client/Services"
)


//func newProd(id int32, pname string) *Services.ProdModel {
//	return &Services.ProdModel{ProdID: id, ProdName: pname}
//}
//
//func defaultProds() (*Services.ProdLIstResponse, error) {
//
//	models := make([]*Services.ProdModel, 0)
//	var i int32
//	for i = 0; i < 1; i++ {
//		models = append(models, newProd(i+200, "prodname"+strconv.Itoa(200+int(i))))
//	}
//	res:=&Services.ProdLIstResponse{}
//	res.Data=models
//	return res, nil
//}


func PanicIfError(err error)  {
	if err!=nil{
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	PanicIfError(ginCtx.BindUri(&prodReq))
	resp,_ := prodService.GetProdDetail(context.Background(),&prodReq)
	ginCtx.JSON(200,gin.H{"data":resp.Data})
}


func GetProdList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
	} else {
		prodRes, _ := prodService.GetProdLIst(context.Background(), &prodReq)

		//熔断器配置
		//configA := hystrix.CommandConfig{
		//	Timeout: 1000,
		//}
		//hystrix.ConfigureCommand("getprods", configA)
		//var prodRes *Services.ProdLIstResponse
		//err := hystrix.Do("getprods", func() error {
		//	prodRes, err = prodService.GetProdLIst(context.Background(), &prodReq)
		//	return err
		//}, func(err error) error {
		//	prodRes,err = defaultProds()
		//	return err
		//})

		ginCtx.JSON(200, gin.H{"data": prodRes.Data})

	}

}
