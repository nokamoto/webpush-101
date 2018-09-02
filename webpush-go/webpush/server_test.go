package main

import (
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"testing"
)

func TestServer_SendPushSubscriptionNotification_empty(t *testing.T) {
	test(t, func(_ string, cli pb.WebpushServiceClient, ctx context.Context) {
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
	test(t, func(url string, cli pb.WebpushServiceClient, ctx context.Context) {
		subscription := &pb.PushSubscription{
			Endpoint: url,
			P256Dh:   fdecode("BOVFfCoBB/2Sn6YZrKytKc1asM+IOXFKz6+T1NLOnrGrRXh/xJEgiJIoFBO9I6twWDAj6OYvhval8jxq8F4K0iM="),
			Auth:     fdecode("LsUmSxGzGt+KcuczkTfFrQ=="),
		}

		req := pb.PushSubscriptionNotification{
			Subscription: []*pb.PushSubscription{subscription},
			Request:      &pb.WebpushRequest{},
		}
		_, err := cli.SendPushSubscriptionNotification(ctx, &req)

		if err != nil {
			t.Fatal(err)
		}
	})
}
