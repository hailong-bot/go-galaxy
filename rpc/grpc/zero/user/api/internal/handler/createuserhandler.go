package handler

import (
	"net/http"

	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/logic"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/svc"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/zero/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
