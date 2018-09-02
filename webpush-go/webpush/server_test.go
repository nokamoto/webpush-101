package main

import (
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServer_SendPushSubscriptionNotification_empty(t *testing.T) {
	test(t, func(cli pb.WebpushServiceClient, ctx context.Context) {
		req := pb.PushSubscriptionNotification{
			Subscription: []*pb.PushSubscription{},
			Request:      &pb.WebpushRequest{},
		}
		_, err := cli.SendPushSubscriptionNotification(ctx, &req)

		if err != nil {
			t.Error(err)
		}
	})
}

func TestServer_SendPushSubscriptionNotification_non_empty(t *testing.T) {
	test(t, func(cli pb.WebpushServiceClient, ctx context.Context) {
		req := pb.PushSubscriptionNotification{
			Subscription: []*pb.PushSubscription{&pb.PushSubscription{}},
			Request:      &pb.WebpushRequest{},
		}
		_, err := cli.SendPushSubscriptionNotification(ctx, &req)

		if status.Convert(err).Code() != codes.Unimplemented {
			t.Errorf("expected %v but actual %v", codes.Unimplemented, status.Convert(err).Code())
		}
	})
}
