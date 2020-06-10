package common

import (
	"crypto/hmac"
	"crypto/sha256"
)

const (
	// HMAC签名算法
	HMAC_SHA256 = "HMAC-SHA256"
)

// HMAC_Sha256 加密
func HMAC_Sha256(data string, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

// RSASHA256 加密
// func RSASHA256(data string, Key string) (string, error) {
// 	h := sha256.New()
// 	h.Write([]byte(data))
// 	hashed := h.Sum(nil)
// 	block, _ := pem.Decode([]byte(Key))
// 	if block == nil {
// 		return "", errors.New("private key error")
// 	}

// 	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	signature, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey, crypto.SHA256, hashed)
// 	if err != nil {
// 		return "", err
// 	}

// 	return base64.StdEncoding.EncodeToString(signature), nil
// }
