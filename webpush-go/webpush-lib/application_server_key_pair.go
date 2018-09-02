package webpushlib

import (
	"crypto/ecdsa"
	"encoding/base64"
	"math/big"
)

// ApplicationServerKeyPair is a key pair for voluntary application server identification (VAPID https://tools.ietf.org/html/rfc8292).
type ApplicationServerKeyPair struct {
	priv *ecdsa.PrivateKey
	pub  string
}

// NewApplicationServerKeyPairFromBase64StdEncodingKeyPair returns a new key pair represented by base64 encoded strings.
func NewApplicationServerKeyPairFromBase64StdEncodingKeyPair(priv, pub string) (*ApplicationServerKeyPair, error) {
	pair := &ApplicationServerKeyPair{}
	if err := pair.SetBase64StdEncodingPrivateKey(priv); err != nil {
		return nil, err
	}
	if err := pair.SetBase64StdEncodingPublicKey(pub); err != nil {
		return nil, err
	}
	return pair, nil
}

// SetBase64StdEncodingPrivateKey set the base64 encoded string as the private key.
func (pair *ApplicationServerKeyPair) SetBase64StdEncodingPrivateKey(key string) error {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}

	d := big.Int{}

	pair.priv = &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: curve},
		D:         d.SetBytes(b),
	}

	return nil
}

// SetBase64StdEncodingPublicKey set the base64 encoded string as the public key.
func (pair *ApplicationServerKeyPair) SetBase64StdEncodingPublicKey(key string) error {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}

	pair.pub = base64.RawURLEncoding.EncodeToString(b)
	return nil
}
