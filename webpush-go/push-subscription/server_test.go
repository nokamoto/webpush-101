package main

import (
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

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
