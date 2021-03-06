// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ProdService.proto

package Services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProdService service

type ProdService interface {
	GetProdLIst(ctx context.Context, in *ProdsRequest, opts ...client.CallOption) (*ProdLIstResponse, error)
	GetProdDetail(ctx context.Context, in *ProdsRequest, opts ...client.CallOption) (*ProdDetailResponse, error)
}

type prodService struct {
	c    client.Client
	name string
}

func NewProdService(name string, c client.Client) ProdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "Services"
	}
	return &prodService{
		c:    c,
		name: name,
	}
}

func (c *prodService) GetProdLIst(ctx context.Context, in *ProdsRequest, opts ...client.CallOption) (*ProdLIstResponse, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdLIst", in)
	out := new(ProdLIstResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodService) GetProdDetail(ctx context.Context, in *ProdsRequest, opts ...client.CallOption) (*ProdDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProdService.GetProdDetail", in)
	out := new(ProdDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProdService service

type ProdServiceHandler interface {
	GetProdLIst(context.Context, *ProdsRequest, *ProdLIstResponse) error
	GetProdDetail(context.Context, *ProdsRequest, *ProdDetailResponse) error
}

func RegisterProdServiceHandler(s server.Server, hdlr ProdServiceHandler, opts ...server.HandlerOption) error {
	type prodService interface {
		GetProdLIst(ctx context.Context, in *ProdsRequest, out *ProdLIstResponse) error
		GetProdDetail(ctx context.Context, in *ProdsRequest, out *ProdDetailResponse) error
	}
	type ProdService struct {
		prodService
	}
	h := &prodServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProdService{h}, opts...))
}

type prodServiceHandler struct {
	ProdServiceHandler
}

func (h *prodServiceHandler) GetProdLIst(ctx context.Context, in *ProdsRequest, out *ProdLIstResponse) error {
	return h.ProdServiceHandler.GetProdLIst(ctx, in, out)
}

func (h *prodServiceHandler) GetProdDetail(ctx context.Context, in *ProdsRequest, out *ProdDetailResponse) error {
	return h.ProdServiceHandler.GetProdDetail(ctx, in, out)
}
