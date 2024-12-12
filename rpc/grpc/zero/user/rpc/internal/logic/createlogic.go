package logic

import (
	"context"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/models"

	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/internal/svc"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		Id: in.Id,
	})
	return &user.GetUserResp{}, nil
}
