package svc

import (
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/config"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
