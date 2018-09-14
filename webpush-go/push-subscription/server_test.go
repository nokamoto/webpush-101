package main

import (
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"testing"
)

func assertCode(err error, t *testing.T, expected codes.Code) {
	if actual := status.Convert(err).Code(); expected != actual {
		t.Errorf("expected %v but actual %v", expected, actual)
	}
}

func assertPushSubscription(t *testing.T, expected []*pb.PushSubscription, actual []*pb.PushSubscription) {
	if len(expected) != len(actual) {
		t.Fatalf("expected %v but actual %v", expected, actual)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i].GetEndpoint() != actual[i].GetEndpoint() {
			t.Fatalf("%d: expected %v but actual %v", i, expected, actual)
		}
	}
}

func get(ctx context.Context, t *testing.T, user *pb.User, cli pb.PushSubscriptionServiceClient) []*pb.PushSubscription {
	stream, err := cli.Get(ctx, user)
	assertCode(err, t, codes.OK)

	res := []*pb.PushSubscription{}

	for {
		s, err := stream.Recv()
		if err == io.EOF {
			break
		}

		assertCode(err, t, codes.OK)

		res = append(res, s)
	}

	return res
}

func TestServer_Subscribe_empty(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := pb.UserSubscription{}

		_, err := cli.Subscribe(ctx, &req)
		assertCode(err, t, codes.OK)
	})
}

func TestServer_Unsubscribe_empty(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := pb.PushSubscription{}

		_, err := cli.Unsubscribe(ctx, &req)
		assertCode(err, t, codes.OK)
	})
}

func TestServer_Get_empty(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		req := &pb.User{}

		actual := get(ctx, t, req, cli)
		if len(actual) != 0 {
			t.Errorf("non empty: %v", actual)
		}
	})
}

func TestServer_integrated(t *testing.T) {
	test(t, func(cli pb.PushSubscriptionServiceClient, ctx context.Context) {
		user := &pb.User{Id: "test"}
		s1 := &pb.PushSubscription{Endpoint: "https://example.com/test/1"}
		s2 := &pb.PushSubscription{Endpoint: "https://example.com/test/2"}
		s3 := &pb.PushSubscription{Endpoint: "https://example.com/test/3"}

		e1 := []*pb.PushSubscription{s1, s2, s3}
		cli.Subscribe(ctx, &pb.UserSubscription{User: user, Subscription: e1})

		assertPushSubscription(t, e1, get(ctx, t, user, cli))

		e2 := []*pb.PushSubscription{s1, s3}
		cli.Unsubscribe(ctx, s2)

		assertPushSubscription(t, e2, get(ctx, t, user, cli))
	})
}
