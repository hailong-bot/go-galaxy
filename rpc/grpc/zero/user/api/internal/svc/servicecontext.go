package svc

import (
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
