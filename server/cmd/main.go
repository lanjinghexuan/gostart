package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/handler"
	_ "server/inits/init"
	pd1 "server/proto/code"
	pb "server/proto/videoUser"
	"server/proto/videoWorks"
	"server/proto/videoWorksComment"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVideoUserServer(s, &handler.VideoUserServer{})
	pd1.RegisterCodeServer(s, &handler.CodeServer{})
	videoWorks.RegisterVideoWorksServer(s, &handler.VideoWorksServer{})
	videoWorksComment.RegisterVideoWorksCommentServer(s, &handler.VideoWorksCommentServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
