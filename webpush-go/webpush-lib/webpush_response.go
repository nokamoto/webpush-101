package webpushlib

import (
	"time"
)

// WebpushResponse is a http response from the push service.
type WebpushResponse struct {
	Status     string
	StatusCode int
	Text       string
	RetryAfter *time.Time
}
