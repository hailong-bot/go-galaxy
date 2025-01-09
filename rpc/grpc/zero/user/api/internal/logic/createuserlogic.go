package logic

import (
	"context"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/userclient"

	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/svc"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// todo: add your logic here and delete this line
	createBody, err := l.svcCtx.UserClient.Create(l.ctx, &userclient.CreateReq{
		Name:  req.Name,
		Phone: req.Mobile,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateUserResp{
		Id: createBody.Id,
	}, nil
}
