package svc

import (
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/models"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserModel(sqlConn),
	}
}
