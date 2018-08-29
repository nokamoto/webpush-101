package webpushlib

import (
	"encoding/binary"
	"io"
	"crypto/sha256"
	"crypto/rand"
	"errors"
	"crypto/elliptic"
	"crypto/cipher"
	"crypto/aes"
	"golang.org/x/crypto/hkdf"
)

type applicationServerKeys struct {
	private []byte
	public []byte
}

var curve = elliptic.P256()

func newApplicationServerKeys() (*applicationServerKeys, error) {
	priv, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	pub := elliptic.Marshal(curve, x, y)
	return &applicationServerKeys{private: priv, public: pub}, nil
}

func newSharedSecret(as *applicationServerKeys, p256dh []byte) ([]byte, error) {
	// ecdh_secret = ECDH(ua_private, as_public)
	x, y := elliptic.Unmarshal(curve, p256dh)
	if x == nil {
		return nil, errors.New("elliptic.Unmarshal failed")
	}
	secret, _ := curve.ScalarMult(x, y, as.private)
	return secret.Bytes(), nil
}

func random(size int) ([]byte, error) {
	bytes := make([]byte, size)
	n, err := io.ReadFull(rand.Reader, bytes)
	if n != len(bytes) {
		return nil, errors.New("ReadFull(rand.Reader) length error")
	}
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newKeyInfo(ua, as []byte) []byte {
	// key_info = "WebPush: info" || 0x00 || ua_public || as_public
	prefix := []byte("WebPush: info\x00")
	bytes := make([]byte, len(prefix) + len(ua) + len(as))
	n := copy(bytes, prefix)
	n += copy(bytes[n:], ua)
	copy(bytes[n:], as)
	return bytes
}

func _hkdf(secret, salt, info []byte, l int) ([]byte, error) {
	h := hkdf.New(sha256.New, secret, salt, info)
	bytes := make([]byte, l)
	n, err := io.ReadFull(h, bytes)
	if n != len(bytes) {
		return nil, errors.New("ReadFull(hkdf) length error")
	}
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newIkm(ua, as, auth, ecdh []byte) ([]byte, error) {
	keyInfo := newKeyInfo(ua, as)

	// PRK_key = HMAC-SHA-256(auth_secret, ecdh_secret)
	// IKM = HMAC-SHA-256(PRK_key, key_info || 0x01)
	return _hkdf(ecdh, auth, keyInfo, 32)
}

// cek_info = "Content-Encoding: aes128gcm" || 0x00
var cekInfo = []byte("Content-Encoding: aes128gcm\x00")

func newCek(salt, ikm []byte) ([]byte, error) {
	// PRK = HMAC-SHA-256(salt, IKM)
	// CEK = HMAC-SHA-256(PRK, cek_info || 0x01)[0..15]
	return _hkdf(ikm, salt, cekInfo, 16)
}

// nonce_info = "Content-Encoding: nonce" || 0x00
var nonceInfo = []byte("Content-Encoding: nonce\x00")

func newNonce(salt, ikm []byte) ([]byte, error) {
	// PRK = HMAC-SHA-256(salt, IKM)
	// NONCE = HMAC-SHA-256(PRK, nonce_info || 0x01)[0..11]
	return _hkdf(ikm, salt, nonceInfo, 12)
}

func newSalt() ([]byte, error) {
	// salt = random(16)
	salt, err := random(16)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func newCiphertext(plaintext string, cek, nonce []byte) ([]byte, error) {
	// https://tools.ietf.org/html/rfc8188#section-2
	b := []byte(plaintext)
	b = append(b, 0x02)

	block, err := aes.NewCipher(cek)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nil, nonce, b, nil), nil
}

func newHeader(salt, keyid []byte) []byte {
	// https://tools.ietf.org/html/rfc8188#section-2.1
	var b []byte

	b = append(b, salt...)

	rs := make([]byte, 4)
	binary.BigEndian.PutUint32(rs, 4096)
	b = append(b, rs...)

	b = append(b, byte(len(keyid)))

	b = append(b, keyid...)

	return b
}

func newContent(header, cihpertext []byte) []byte {
	return append(header, cihpertext...)
}

func encrypt(as *applicationServerKeys, salt []byte, body string, p256dh, auth []byte) ([]byte, error) {
	// https://tools.ietf.org/html/rfc8291
	sharedSecret, err := newSharedSecret(as, p256dh)
	if err != nil {
		return nil, err
	}

	ikm, err := newIkm(p256dh, as.public, auth, sharedSecret)
	if err != nil {
		return nil, err
	}

	cek, err := newCek(salt, ikm)
	if err != nil {
		return nil, err
	}

	nonce, err := newNonce(salt, ikm)
	if err != nil {
		return nil, err
	}

	header := newHeader(salt, as.public)

	ciphertext, err := newCiphertext(body, cek, nonce)
	if err != nil {
		return nil, err
	}

	return newContent(header, ciphertext), nil
}
