package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"github.com/nokamoto/webpush-101/webpush-go/webpush-lib"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
)

type server struct {
	client *webpushlib.PushServiceClient
}

func (s *server) SendPushSubscriptionNotification(_ context.Context, req *pb.PushSubscriptionNotification) (*empty.Empty, error) {
	log.Printf("SendPushSubscriptionNotification(%v)\n", req)
	for _, subscription := range req.GetSubscription() {
		res, err := s.client.Send(subscription, req.GetRequest())

		if err != nil {
			return nil, grpc.Errorf(codes.Internal, err.Error())
		}

		if res.StatusCode != 201 {
			return nil, grpc.Errorf(codes.Internal, fmt.Sprintf("%d %s", res.StatusCode, res.Status))
		}
	}
	return &empty.Empty{}, nil
}
