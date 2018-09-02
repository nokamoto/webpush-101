package webpushlib

import (
	"errors"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"net/http"
)

// PushServiceClient is a http client for webpush.
type PushServiceClient struct {
	KeyPair *ApplicationServerKeyPair
	Client  *http.Client
}

// Send sends the webpush request with the push subscription.
func (c *PushServiceClient) Send(subscription *pb.PushSubscription, request *pb.WebpushRequest) (*http.Response, error) {
	return nil, errors.New("not implemented yet")
}
