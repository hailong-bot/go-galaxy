package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("log begin")
	resp, err = handler(ctx, req)
	fmt.Println("log end")
	return
}
