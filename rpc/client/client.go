package main

import (
	"log"
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

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var (
		req = GetUserReq{
			Id: "2",
		}
		res GetUserResp
	)
	err = client.Call("UserServer.GetUser", req, &res)
	if err != nil {
		log.Fatal("请求异常")
		return
	}
	log.Println(res)
}
