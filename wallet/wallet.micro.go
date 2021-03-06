// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: wallet/wallet.proto

/*
Package go_micro_srv_wallet is a generated protocol buffer package.

It is generated from these files:
	wallet/wallet.proto

It has these top-level messages:
	CreateWalletRequest
	CreateWalletResponse
*/
package go_micro_srv_wallet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for WalletService service

type WalletServiceClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...client.CallOption) (*CreateWalletResponse, error)
}

type walletServiceClient struct {
	c           client.Client
	serviceName string
}

func NewWalletServiceClient(serviceName string, c client.Client) WalletServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.wallet"
	}
	return &walletServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...client.CallOption) (*CreateWalletResponse, error) {
	req := c.c.NewRequest(c.serviceName, "WalletService.CreateWallet", in)
	out := new(CreateWalletResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WalletService service

type WalletServiceHandler interface {
	CreateWallet(context.Context, *CreateWalletRequest, *CreateWalletResponse) error
}

func RegisterWalletServiceHandler(s server.Server, hdlr WalletServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&WalletService{hdlr}, opts...))
}

type WalletService struct {
	WalletServiceHandler
}

func (h *WalletService) CreateWallet(ctx context.Context, in *CreateWalletRequest, out *CreateWalletResponse) error {
	return h.WalletServiceHandler.CreateWallet(ctx, in, out)
}
