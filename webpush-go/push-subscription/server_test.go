package main

import (
	"fmt"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"testing"
	"time"
)

func withServer(t *testing.T, f func(int)) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	defer s.GracefulStop()

	pb.RegisterPushSubscriptionServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	f(port)
}

func withClient(t *testing.T, port int, f func(pb.PushSubscriptionServiceClient, context.Context)) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPushSubscriptionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	f(client, ctx)
}

func test(t *testing.T, f func(pb.PushSubscriptionServiceClient, context.Context)) {
	withServer(t, func(port int) {
		withClient(t, port, f)
	})
}

func TestServer_Subscribe(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := pb.UserSubscription{}
		_, err := cli.Subscribe(ctx, &req)
		if status.Convert(err).Code() != codes.Unimplemented {
			t.Errorf("expected %v but actual %v", codes.Unimplemented, status.Convert(err).Code())
		}
	})
}

func TestServer_Unsubscribe(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := pb.PushSubscription{}
		_, err := cli.Unsubscribe(ctx, &req)
		if status.Convert(err).Code() != codes.Unimplemented {
			t.Errorf("expected %v but actual %v", codes.Unimplemented, status.Convert(err).Code())
		}
	})
}

func TestServer_Get(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := pb.User{}
		res, err := cli.Get(ctx, &req)
		if status.Convert(err).Code() != codes.OK {
			t.Fatalf("expected %v but actual %v", codes.OK, status.Convert(err).Code())
		}

		_, err = res.Recv()
		if status.Convert(err).Code() != codes.Unimplemented {
			t.Errorf("expected %v but actual %v", codes.Unimplemented, status.Convert(err).Code())
		}
	})
}
