package svc

import (
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/config"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/middleware"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/models"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	UserClient        userclient.User
	UserModel         models.UserModel
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:            c,
		UserClient:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserModel:         models.NewUserModel(sqlConn),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
