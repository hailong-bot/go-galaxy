package logic

import (
	"context"
	"database/sql"
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

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	insertPO, err := l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		Id:     in.Id,
		Mobile: in.Phone,
		Name: sql.NullString{
			String: in.Name,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}
	id, err := insertPO.LastInsertId()
	if err != nil {
		return nil, err

	}
	return &user.CreateResp{
		Id: id,
	}, nil
}
