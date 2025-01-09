package logic

import (
	"context"
	"strconv"

	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/internal/svc"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	userID, _ := strconv.ParseInt(in.Id, 10, 64)
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	result := &user.GetUserResp{
		Id:    in.Id,
		Name:  one.Name.String,
		Phone: "110",
	}

	return result, nil
}
