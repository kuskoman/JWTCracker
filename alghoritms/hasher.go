package alghoritms

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

type Hasher interface {
	Sign(body, secret string) string
}

type HS256Hasher struct{}

func (h *HS256Hasher) Sign(body, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash)
}

func GetHasher(alg string) Hasher {
	var h Hasher

	switch alg {
	case "HS256":
		h = &HS256Hasher{}
	default:
		log.Fatal("This alghoritm is not supported yet")
	}
	return h
}