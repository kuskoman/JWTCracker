package alghoritms

import (
	"crypto/hmac"
	"encoding/base64"
	"hash"
)

type SHAHasher struct {
	Alg func() hash.Hash
}

func (h *SHAHasher) Sign(body, secret string) string {
	mac := hmac.New(h.Alg, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(hash)
}
