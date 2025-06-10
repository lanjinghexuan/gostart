package server

import (
	pb "api/proto/Code"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
)

type CodeHandle func(ctx context.Context, client pb.CodeClient) (interface{}, error)

func CodeServer(c context.Context, client CodeHandle) (interface{}, error) {
	flag.Parse()
	coon, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	c1 := pb.NewCodeClient(coon)
	return client(c, c1)
}

func SendCode(c context.Context, req *pb.SendCodeReq) (*pb.SendCodeRes, error) {
	res, err := CodeServer(c, func(ctx context.Context, client pb.CodeClient) (interface{}, error) {
		res, err := client.SendCode(ctx, req)
		return res, err
	})
	return res.(*pb.SendCodeRes), err
}
