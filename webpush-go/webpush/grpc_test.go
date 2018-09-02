package main

import (
	"fmt"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"github.com/nokamoto/webpush-101/webpush-go/webpush-lib"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func withServer(t *testing.T, f func(int, string)) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	defer s.GracefulStop()

	mock := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(201)
	}))
	defer mock.Close()

	pair, err := webpushlib.NewApplicationServerKeyPairFromBase64StdEncodingKeyPair(
		"AJFotoB4FS7IX6tbm5t0SGyISTQ6l54mMzpfYipdOD+N",
		"BNuvjW90TpDawYyxhvK79QVyNEplaSQZOWo1CwXDmWwfya6qnyBvIx3tFvKEBetExvil4rNNRL0/ZR2WLjGEAbQ=")
	if err != nil {
		t.Fatal(err)
	}

	client := &webpushlib.PushServiceClient{KeyPair: pair, Client: mock.Client()}

	pb.RegisterWebpushServiceServer(s, &server{client: client})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	f(port, mock.URL)
}

func withClient(t *testing.T, port int, url string, f func(string, pb.WebpushServiceClient, context.Context)) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWebpushServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	f(url, client, ctx)
}

func test(t *testing.T, f func(string, pb.WebpushServiceClient, context.Context)) {
	withServer(t, func(port int, url string) {
		withClient(t, port, url, f)
	})
}
