package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func ErrInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("err begin")
	resp, err = handler(ctx, req)
	fmt.Println("err end")
	return
}
