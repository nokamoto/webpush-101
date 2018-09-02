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

		// https://tools.ietf.org/html/rfc8030
		// A push service MUST return a 400 (Bad Request) status code for requests that contain an invalid receipt subscription.
		//
		// A push service MUST return a 400 (Bad Request) status code in response to requests that omit the TTL header field.
		//
		// https://tools.ietf.org/html/rfc8292
		// Though a push service is not obligated to check either parameter for every push message, a push service SHOULD reject push messages that have identical values for these parameters with a 400 (Bad Request) status code.
		// (* Note that this condition may prefer to return FailedPrecondition rather than return InvalidArgument.)
		if res.StatusCode == 400 {
			return nil, grpc.Errorf(codes.InvalidArgument, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		// https://tools.ietf.org/html/rfc8292
		// A push service MAY reject a request with a 403 (Forbidden) status code [RFC7231] if the JWT signature or its claims are invalid.
		if res.StatusCode == 403 {
			return nil, grpc.Errorf(codes.Unauthenticated, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		// https://tools.ietf.org/html/rfc8030
		// A push service MUST return a 404 (Not Found) status code if an application server attempts to send a push message to an expired push message subscription.
		if res.StatusCode == 404 {
			return nil, grpc.Errorf(codes.InvalidArgument, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		// https://tools.ietf.org/html/rfc8030
		// To limit the size of messages, the push service MAY return a 413 (Payload Too Large) status code [RFC7231] in response to requests that include an entity body that is too large
		if res.StatusCode == 413 {
			return nil, grpc.Errorf(codes.InvalidArgument, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		// https://tools.ietf.org/html/rfc8030
		// If a push service wishes to limit the number of receipt subscriptions that it maintains, it MAY return a 429 (Too Many Requests) status code [RFC6585] to reject receipt requests that omit a receipt subscription.
		// The push service SHOULD also include a Retry-After header [RFC7231] to indicate how long the application server is requested to wait before it makes another request to the push resource.
		if res.StatusCode == 429 {
			return nil, grpc.Errorf(codes.ResourceExhausted, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		if res.StatusCode/100 == 4 {
			return nil, grpc.Errorf(codes.InvalidArgument, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		if res.StatusCode/100 == 5 {
			return nil, grpc.Errorf(codes.Unavailable, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}
	}
	return &empty.Empty{}, nil
}
