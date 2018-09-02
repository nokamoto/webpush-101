package webpushlib

import (
	"net/http"
	"testing"
	"time"
)

func TestParseRetryAfter_none(t *testing.T) {
	header := &http.Header{}

	after, err := parseRetryAfter(time.Now(), header)
	if err != nil {
		t.Fatal(err)
	}
	if after != nil {
		t.Fatalf("retry after expected nil but actual %v", after)
	}
}

func TestParseRetryAfter_delay_seconds(t *testing.T) {
	now := time.Now()
	expected := now.Add(time.Duration(60) * time.Second)

	header := &http.Header{}
	header.Add("Retry-After", "60")

	after, err := parseRetryAfter(now, header)
	if err != nil {
		t.Fatal(err)
	}
	if after.Unix() != expected.Unix() {
		t.Fatalf("retry after %v but actual %v", expected, after)
	}
}

func TestParseRetryAfter_http_date(t *testing.T) {
	expected := time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC)

	header := &http.Header{}
	header.Add("Retry-After", "Fri, 31 Dec 1999 23:59:59 GMT")

	after, err := parseRetryAfter(time.Now(), header)
	if err != nil {
		t.Fatal(err)
	}
	if after.Unix() != expected.Unix() {
		t.Fatalf("retry after %v but actual %v", expected, after)
	}
}
