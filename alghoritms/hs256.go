package alghoritms

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type HS256Hasher struct{}

func (h *HS256Hasher) Sign(body, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash)
}
