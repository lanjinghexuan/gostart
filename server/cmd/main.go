package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/handler"
	"server/inits/config"
	"server/inits/mysql"
	"server/inits/nacos"
	"server/inits/redis"
	pb "server/proto/user"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	config.GetConfig()
	nacos.NacosConfig()
	mysql.MysqlInit()
	redis.RedisInit()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &handler.UserServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
