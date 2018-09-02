package webpushlib

import (
	"bytes"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"io/ioutil"
	"net/http"
	"strconv"
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

func parseRetryAfter(now time.Time, h *http.Header) (*time.Time, error) {
	// https://tools.ietf.org/html/rfc7231#section-7.1.3
	retryAfter := h.Get("Retry-After")
	if retryAfter == "" {
		return nil, nil
	}

	if seconds, err := strconv.ParseUint(retryAfter, 10, 31); err == nil {
		t := now.Add(time.Duration(seconds) * time.Second)
		return &t, nil
	}

	t, err := time.Parse(time.RFC1123, retryAfter)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// Send sends the webpush request with the push subscription.
func (c *PushServiceClient) Send(subscription *pb.PushSubscription, request *pb.WebpushRequest) (*WebpushResponse, error) {
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
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	retryAfter, err := parseRetryAfter(time.Now(), &res.Header)
	if err != nil {
		// todo
	}

	wr := &WebpushResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Text:       string(b),
		RetryAfter: retryAfter,
	}

	return wr, nil
}
