package alghoritms

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
)

type HS384Hasher struct{}

func (h *HS384Hasher) Sign(body, secret string) string {
	mac := hmac.New(sha512.New384, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash)
}
