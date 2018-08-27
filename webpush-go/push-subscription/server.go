package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
)

type server struct{}

func (s *server) Subscribe(_ context.Context, req *pb.UserSubscription) (*empty.Empty, error) {
	log.Printf("Subscribe(%v)\n", req)
	return nil, grpc.Errorf(codes.Unimplemented, "not implemented yet")
}

func (s *server) Unsubscribe(_ context.Context, req *pb.PushSubscription) (*empty.Empty, error) {
	log.Printf("Unsubscribe(%v)", req)
	return nil, grpc.Errorf(codes.Unimplemented, "not implemented yet")
}

func (s *server) Get(req *pb.User, _ pb.PushSubscriptionService_GetServer) error {
	log.Printf("Get(%v)", req)
	return grpc.Errorf(codes.Unimplemented, "not implemented yet")
}
