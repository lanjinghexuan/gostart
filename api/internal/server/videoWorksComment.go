package server

import (
	pb "api/proto/videoWorksComment"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

type VideoWorksCommentClientHand func(c context.Context, client pb.VideoWorksCommentClient) (interface{}, error)

func VideoWorksCommentClient(c context.Context, client VideoWorksCommentClientHand) (interface{}, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := pb.NewVideoWorksCommentClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return client(ctx, c1)
}

func GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentRes, error) {
	res, err := VideoWorksCommentClient(ctx, func(c context.Context, client pb.VideoWorksCommentClient) (interface{}, error) {
		res, err := client.GetComment(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetCommentRes), nil
}

func GetCommentReplyList(ctx context.Context, req *pb.GetCommentReplyListReq) (*pb.GetCommentReplyListRes, error) {
	res, err := VideoWorksCommentClient(ctx, func(c context.Context, client pb.VideoWorksCommentClient) (interface{}, error) {
		res, err := client.GetCommentReplyList(ctx, req)
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetCommentReplyListRes), nil
}
