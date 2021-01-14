package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro/Models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func callAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("prodservise", "/v1/prods",
		Models.ProdsRequest{Size: 4})
	var rsp Models.ProdLIstResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())
}

func callAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {

	consulReg := consul.NewRegistry(
		registry.Addrs("techad.top:8500"),
	)

	// callAPI调用方法
	{
		getService, err := consulReg.GetService("prodservise")
		if err != nil {
			log.Fatal(err)
		}

		next := selector.Random(getService)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}

		callRes, err := callAPI(node.Address, "/v1/prods", "POST")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(callRes)
		fmt.Println(node.Address, node.Id, node.Metadata)

	}

	time.Sleep(time.Second)

	// callAPI2 调用方法
	{
		mySelector := selector.NewSelector(
			selector.Registry(consulReg),
			selector.SetStrategy(selector.Random),
		)

		callAPI2(mySelector)
	}

}
