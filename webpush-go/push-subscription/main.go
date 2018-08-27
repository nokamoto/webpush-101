package main

import (
	"flag"
	"fmt"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port = flag.Int("p", 8000, "gRPC server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listen %v port", *port)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	pb.RegisterPushSubscriptionServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
