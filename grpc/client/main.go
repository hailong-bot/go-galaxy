package main

import (
	"context"
	"log"

	"github.com/hailong-bot/go-galaxy/grpc/protoc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *grpc.ClientConn) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	c := user.NewUserClient(client)
	res, err := c.GetUser(context.Background(), &user.GetUserReq{Id: "1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
