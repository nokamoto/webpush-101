package webpushlib

import (
	"testing"
)

// https://tools.ietf.org/html/draft-ietf-webpush-encryption-09#section-5
// https://tools.ietf.org/html/draft-ietf-webpush-encryption-09#appendix-A
var body = "When I grow up, I want to be a watermelon"

var asPrivate = fdecode("yfWPiYE-n46HLnH0KqZOF1fJJU3MYrct3AELtAQ-oRw")

var asPublic = fdecode("BP4z9KsN6nGRTbVYI_c7VJSPQTBtkgcy27mlmlMoZIIgDll6e3vCYLocInmYWAmS6TlzAC8wEqKK6PBru3jl7A8")

var salt = fdecode("DGv6ra1nlYgDCS1FRnbzlw")

var uaPublic = fdecode("BCVxsr7N_eNgVRqvHtD0zTZsEc6-VV-JvLexhqUzORcxaOzi6-AYWXvTBHm4bjyPjs7Vd8pZGH6SRpkNtoIAiw4")

var authSecret = fdecode("BTBZMqHH6r4Tts7J_aSIgg")

func as() *applicationServerKeys {
	return &applicationServerKeys{private: asPrivate, public: asPublic}
}

func TestEncrypt_body(t *testing.T) {
	expected := "V2hlbiBJIGdyb3cgdXAsIEkgd2FudCB0byBiZSBhIHdhdGVybWVsb24"

	if s := encode([]byte(body)); s != expected {
		t.Fatalf("plaintext expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_ecdh_secret(t *testing.T) {
	expected := "kyrL1jIIOHEzg3sM2ZWRHDRB62YACZhhSlknJ672kSs"

	actual, err := newSharedSecret(as(), uaPublic)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(actual); s != expected {
		t.Fatalf("ecdh_secret expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_key_info(t *testing.T) {
	expected := "V2ViUHVzaDogaW5mbwAEJXGyvs3942BVGq8e0PTNNmwRzr5VX4m8t7GGpTM5FzFo7OLr4BhZe9MEebhuPI-OztV3ylkYfpJGmQ22ggCLDgT-M_SrDepxkU21WCP3O1SUj0EwbZIHMtu5pZpTKGSCIA5Zent7wmC6HCJ5mFgJkuk5cwAvMBKiiujwa7t45ewP"

	actual := newKeyInfo(uaPublic, as().public)
	if s := encode(actual); s != expected {
		t.Fatalf("key_info expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_ikm(t *testing.T) {
	expected := "S4lYMb_L0FxCeq0WhDx813KgSYqU26kOyzWUdsXYyrg"

	sharedSecret, _ := newSharedSecret(as(), uaPublic)

	ikm, err := newIkm(uaPublic, as().public, authSecret, sharedSecret)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(ikm); s != expected {
		t.Fatalf("ikm expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_cek_info(t *testing.T) {
	expected := "Q29udGVudC1FbmNvZGluZzogYWVzMTI4Z2NtAA"

	if s := encode(cekInfo); s != expected {
		t.Fatalf("cek_info expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_cek(t *testing.T) {
	expected := "oIhVW04MRdy2XN9CiKLxTg"

	sharedSecret, _ := newSharedSecret(as(), uaPublic)
	ikm, _ := newIkm(uaPublic, as().public, authSecret, sharedSecret)

	cek, err := newCek(salt, ikm)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(cek); s != expected {
		t.Fatalf("cek expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_nonce_info(t *testing.T) {
	expected := "Q29udGVudC1FbmNvZGluZzogbm9uY2UA"

	if s := encode(nonceInfo); s != expected {
		t.Fatalf("nonce_info expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_nonce(t *testing.T) {
	expected := "4h_95klXJ5E_qnoN"

	sharedSecret, _ := newSharedSecret(as(), uaPublic)
	ikm, _ := newIkm(uaPublic, as().public, authSecret, sharedSecret)
	
	nonce, err := newNonce(salt, ikm)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(nonce); s != expected {
		t.Fatalf("nonce expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_ciphertext(t *testing.T) {
	expected := "8pfeW0KbunFT06SuDKoJH9Ql87S1QUrdirN6GcG7sFz1y1sqLgVi1VhjVkHsUoEsbI_0LpXMuGvnzQ"

	sharedSecret, _ := newSharedSecret(as(), uaPublic)
	ikm, _ := newIkm(uaPublic, as().public, authSecret, sharedSecret)
	cek, _ := newCek(salt, ikm)
	nonce, _ := newNonce(salt, ikm)

	ciphertext, err := newCiphertext(body, cek, nonce)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(ciphertext); s != expected {
		t.Fatalf("ciphertext expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_header(t *testing.T) {
	expected := "DGv6ra1nlYgDCS1FRnbzlwAAEABBBP4z9KsN6nGRTbVYI_c7VJSPQTBtkgcy27mlmlMoZIIgDll6e3vCYLocInmYWAmS6TlzAC8wEqKK6PBru3jl7A8"

	header := newHeader(salt, as().public)
	if s := encode(header); s != expected {
		t.Fatalf("header expected %s but actual %s", expected, s)
	}
}

func TestEncrypt_content(t *testing.T) {
	expected := "DGv6ra1nlYgDCS1FRnbzlwAAEABBBP4z9KsN6nGRTbVYI_c7VJSPQTBtkgcy27mlmlMoZIIgDll6e3vCYLocInmYWAmS6TlzAC8wEqKK6PBru3jl7A_yl95bQpu6cVPTpK4Mqgkf1CXztLVBSt2Ks3oZwbuwXPXLWyouBWLVWGNWQexSgSxsj_Qulcy4a-fN"

	sharedSecret, _ := newSharedSecret(as(), uaPublic)
	ikm, _ := newIkm(uaPublic, as().public, authSecret, sharedSecret)
	cek, _ := newCek(salt, ikm)
	nonce, _ := newNonce(salt, ikm)
	header := newHeader(salt, as().public)
	ciphertext, _ := newCiphertext(body, cek, nonce)

	content := newContent(header, ciphertext)
	if s := encode(content); s != expected {
		t.Fatalf("content expected %s but actual %s", expected, s)
	}
}

func TestEncrypt(t *testing.T) {
	expected := "DGv6ra1nlYgDCS1FRnbzlwAAEABBBP4z9KsN6nGRTbVYI_c7VJSPQTBtkgcy27mlmlMoZIIgDll6e3vCYLocInmYWAmS6TlzAC8wEqKK6PBru3jl7A_yl95bQpu6cVPTpK4Mqgkf1CXztLVBSt2Ks3oZwbuwXPXLWyouBWLVWGNWQexSgSxsj_Qulcy4a-fN"

	encrypted, err := encrypt(as(), salt, body, uaPublic, authSecret)
	if err != nil {
		t.Fatal(err)
	}
	if s := encode(encrypted); s != expected {
		t.Fatalf("encrypted body expected %s but actual %s", expected, s)	
	}
}
