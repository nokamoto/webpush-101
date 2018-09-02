package main

import (
	"fmt"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	pb.RegisterWebpushServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	f(port)
}

func withClient(t *testing.T, port int, f func(pb.WebpushServiceClient, context.Context)) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWebpushServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	f(client, ctx)
}

func test(t *testing.T, f func(pb.WebpushServiceClient, context.Context)) {
	withServer(t, func(port int) {
		withClient(t, port, f)
	})
}
