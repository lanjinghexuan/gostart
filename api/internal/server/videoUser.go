package server

import (
	pb "api/proto/videoUser"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

type VideoUserClientHand func(c context.Context, client pb.VideoUserClient) (interface{}, error)

func VideoUserClient(c context.Context, client VideoUserClientHand) (interface{}, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := pb.NewVideoUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return client(ctx, c1)
}

func Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	res, err := VideoUserClient(ctx, func(c context.Context, client pb.VideoUserClient) (interface{}, error) {
		res, err := client.Login(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginRes), nil
}

func GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRes, error) {
	res, err := VideoUserClient(ctx, func(c context.Context, client pb.VideoUserClient) (interface{}, error) {
		res, err := client.GetUserInfo(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUserInfoRes), nil
}

func Like(ctx context.Context, req *pb.LikeReq) (*pb.LikeRes, error) {
	res, err := VideoUserClient(ctx, func(c context.Context, client pb.VideoUserClient) (interface{}, error) {
		res, err := client.Like(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.LikeRes), nil
}

func LikeVideo(ctx context.Context, req *pb.LikeVideoReq) (*pb.LikeVideoRes, error) {
	res, err := VideoUserClient(ctx, func(c context.Context, client pb.VideoUserClient) (interface{}, error) {
		res, err := client.LikeVideo(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.LikeVideoRes), nil
}
