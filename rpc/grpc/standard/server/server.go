package main

import (
	"context"
	"errors"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/standard/protoc/user"
	"github.com/hailong-bot/go-galaxy/rpc/grpc/standard/server/interceptor"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (*UserService) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	log.Println("GetUser 调用")
	if u, ok := users[req.Id]; ok {
		resp := &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return resp, nil
	}
	return nil, errors.New("用户未找到")
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.LogInterceptor, interceptor.ErrInterceptor))
	user.RegisterUserServer(s, new(UserService))
	log.Println("启动成功")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
