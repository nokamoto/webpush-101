package webpushlib

import (
	"encoding/base64"
)

func encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func fdecode(s string) []byte {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
