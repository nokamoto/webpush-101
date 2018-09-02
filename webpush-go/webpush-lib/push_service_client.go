package webpushlib

import (
	"bytes"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"net/http"
	"time"
)

// PushServiceClient is a http client for webpush.
type PushServiceClient struct {
	KeyPair *ApplicationServerKeyPair
	Client  *http.Client
}

func (c *PushServiceClient) encrypt(subscription *pb.PushSubscription, request *pb.WebpushRequest) ([]byte, error) {
	as, err := newApplicationServerKeys()
	if err != nil {
		return nil, err
	}

	salt, err := random(16)
	if err != nil {
		return nil, err
	}

	return encrypt(as, salt, request.Content, subscription.P256Dh, subscription.Auth)
}

// Send sends the webpush request with the push subscription.
func (c *PushServiceClient) Send(subscription *pb.PushSubscription, request *pb.WebpushRequest) (*http.Response, error) {
	content, err := c.encrypt(subscription, request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", subscription.Endpoint, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	req.Header.Add("TTL", "30")
	req.Header.Add("Content-Encoding", "aes128gcm")

	subject := "mailto:nokamoto.engr@gmail.com"
	expiry := time.Now().Add(12 * time.Hour).Unix()
	addAuthorizationHeader(req, subscription.Endpoint, subject, expiry, c.KeyPair)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()

	return res, nil
}
