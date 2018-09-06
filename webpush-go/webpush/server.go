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
		// (* It may prefer to return InvalidArgument rather than return Unknown.)
		//
		// A push service MUST return a 400 (Bad Request) status code in response to requests that omit the TTL header field.
		// (* It may prefer to return FailedPrecondition rather than return Unknown.)
		//
		// https://tools.ietf.org/html/rfc8292
		// Though a push service is not obligated to check either parameter for every push message, a push service SHOULD reject push messages that have identical values for these parameters with a 400 (Bad Request) status code.
		// (* It may prefer to return FailedPrecondition rather than return Unknown.)
		if res.StatusCode == 400 {
			return nil, grpc.Errorf(codes.Unknown, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
		}

		// https://tools.ietf.org/html/rfc8292
		// A 401 (Unauthorized) status code might be used if the authentication is absent;
		if res.StatusCode == 401 {
			return nil, grpc.Errorf(codes.FailedPrecondition, fmt.Sprintf("%d %s %s", res.StatusCode, res.Status, res.Text))
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
		// The push service MAY cease to retry delivery of the message prior to its advertised expiration due to scenarios such as an unresponsive user agent or operational constraints.
		// If the application has requested a delivery receipt, then the push service MUST return a 410 (Gone) response to the application server monitoring the receipt subscription.
		//
		// If the user agent fails to acknowledge the receipt of the push message and the push service ceases to retry delivery of the message prior to its advertised expiration, then the push service MUST push a failure response with a status code of 410 (Gone).
		if res.StatusCode == 410 {
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
