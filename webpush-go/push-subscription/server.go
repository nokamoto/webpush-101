package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"log"
	"sync"
)

type server struct {
	inmemory map[string][]*pb.PushSubscription
	mu       sync.Mutex
}

func newServer() *server {
	return &server{inmemory: map[string][]*pb.PushSubscription{}}
}

func (s *server) Subscribe(_ context.Context, req *pb.UserSubscription) (*empty.Empty, error) {
	log.Printf("Subscribe(%v)\n", req)

	s.mu.Lock()
	defer s.mu.Unlock()

	key := req.GetUser().GetId()

	subscription, ok := s.inmemory[key]
	if !ok {
		s.inmemory[key] = req.GetSubscription()
		return &empty.Empty{}, nil
	}

	res := make([]*pb.PushSubscription, len(subscription))
	copy(res, subscription)

	for _, x := range req.GetSubscription() {
		found := false
		for _, y := range subscription {
			found = found || x.GetEndpoint() == y.GetEndpoint()
		}
		if !found {
			res = append(res, x)
		}
	}

	s.inmemory[key] = res

	return &empty.Empty{}, nil
}

func (s *server) Unsubscribe(_ context.Context, req *pb.PushSubscription) (*empty.Empty, error) {
	log.Printf("Unsubscribe(%v)", req)

	s.mu.Lock()
	defer s.mu.Unlock()

	inmemory := map[string][]*pb.PushSubscription{}

	for k, v := range s.inmemory {
		inmemory[k] = []*pb.PushSubscription{}
		for _, s := range v {
			if s.GetEndpoint() != req.GetEndpoint() {
				inmemory[k] = append(inmemory[k], s)
			}
		}
	}

	s.inmemory = inmemory

	return &empty.Empty{}, nil
}

func (s *server) Get(req *pb.User, stream pb.PushSubscriptionService_GetServer) error {
	log.Printf("Get(%v)", req)

	s.mu.Lock()
	defer s.mu.Unlock()

	key := req.GetId()

	subscription, ok := s.inmemory[key]
	if ok {
		for _, s := range subscription {
			err := stream.Send(s)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
