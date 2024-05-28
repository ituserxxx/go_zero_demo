package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero_demo/http/internal/config"
	"go_zero_demo/rpc/hello/helloclient"
)

type ServiceContext struct {
	Config   config.Config
	HelloRpc helloclient.Hello
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		HelloRpc: helloclient.NewHello(zrpc.MustNewClient(c.HelloRpc)),
	}
}
