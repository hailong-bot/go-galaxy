package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type (
	GetUserReq struct {
		Id string `json:"id"`
	}

	GetUserResp struct {
		Id    string
		Name  string
		Phone string
	}
)

type UserServer struct{}

func (u *UserServer) GetUser(req GetUserReq, resp *GetUserResp) error {
	if u, ok := users[req.Id]; ok {
		*resp = GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}
	return errors.New("未找到用户")
}

func main() {
	userServer := new(UserServer)
	rpc.Register(userServer)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("服务器启动成功")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("监听客户端连接失败")
			continue
		}
		go rpc.ServeConn(conn)
	}
}
