package webpushlib

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/url"
)

func newJwt(endpoint, subject string, expiry int64, priv *ecdsa.PrivateKey) (string, error) {
	// https://tools.ietf.org/html/rfc8292
	origin, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	audience := fmt.Sprintf("%s://%s", origin.Scheme, origin.Host)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"aud": audience,
		"exp": expiry,
		"sub": subject,
	})

	return token.SignedString(priv)
}

func addAuthorizationHeader(req *http.Request, endpoint, subject string, expiry int64, pair *ApplicationServerKeyPair) error {
	// https://tools.ietf.org/html/rfc8292
	t, err := newJwt(endpoint, subject, expiry, pair.priv)
	if err != nil {
		return err
	}

	k := pair.pub

	req.Header.Add("Authorization", fmt.Sprintf("vapid t=%s,k=%s", t, k))

	return nil
}
