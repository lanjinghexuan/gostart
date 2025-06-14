package server

import (
	pb "api/proto/videoWorks"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

type VideoWorksClientHand func(c context.Context, client pb.VideoWorksClient) (interface{}, error)

func VideoWorksClient(c context.Context, client VideoWorksClientHand) (interface{}, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := pb.NewVideoWorksClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return client(ctx, c1)
}

func GetVideoWorks(ctx context.Context, req *pb.GetVideoWorksReq) (*pb.GetVideoWorksRes, error) {
	res, err := VideoWorksClient(ctx, func(c context.Context, client pb.VideoWorksClient) (interface{}, error) {
		res, err := client.GetVideoWorks(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetVideoWorksRes), nil
}

func AddVideoWorks(ctx context.Context, req *pb.AddVideoWorksReq) (*pb.AddVideoWorksRes, error) {
	res, err := VideoWorksClient(ctx, func(c context.Context, client pb.VideoWorksClient) (interface{}, error) {
		res, err := client.AddVideoWorks(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.AddVideoWorksRes), nil
}
