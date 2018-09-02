package main

import (
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServer_SendPushSubscriptionNotification_empty(t *testing.T) {
	test(t, 201, func(url string, cli pb.WebpushServiceClient, ctx context.Context) {
		req := &pb.PushSubscriptionNotification{}
		_, err := cli.SendPushSubscriptionNotification(ctx, req)

		if err != nil {
			t.Error(err)
		}
	})
}

func TestServer_SendPushSubscriptionNotification_created(t *testing.T) {
	test(t, 201, func(url string, cli pb.WebpushServiceClient, ctx context.Context) {
		req := testNotification(url)
		_, err := cli.SendPushSubscriptionNotification(ctx, req)

		if err != nil {
			t.Fatal(err)
		}
	})
}

func checkStatusCode(mockStatus int, grpcStatus codes.Code, t *testing.T) {
	test(t, mockStatus, func(url string, cli pb.WebpushServiceClient, ctx context.Context) {
		req := testNotification(url)
		_, err := cli.SendPushSubscriptionNotification(ctx, req)

		if status.Convert(err).Code() != grpcStatus {
			t.Fatalf("error expected %v but actual %v", grpcStatus, err)
		}
	})
}

func TestServer_SendPushSubscriptionNotification_bad_request(t *testing.T) {
	checkStatusCode(400, codes.InvalidArgument, t)
}

func TestServer_SendPushSubscriptionNotification_forbidden(t *testing.T) {
	checkStatusCode(403, codes.Unauthenticated, t)
}

func TestServer_SendPushSubscriptionNotification_not_found(t *testing.T) {
	checkStatusCode(404, codes.InvalidArgument, t)
}

func TestServer_SendPushSubscriptionNotification_payload_too_large(t *testing.T) {
	checkStatusCode(413, codes.InvalidArgument, t)
}

func TestServer_SendPushSubscriptionNotification_too_many_request(t *testing.T) {
	checkStatusCode(429, codes.ResourceExhausted, t)
}

func TestServer_SendPushSubscriptionNotification_internal_server_error(t *testing.T) {
	checkStatusCode(500, codes.Unavailable, t)
}
