package hashing

import (
	"testing"

	"github.com/kuskoman/JWTCracker/utils"
)

func TestHS256Hasher(t *testing.T) {
	payload := "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
	secret := "your-256-bit-secret"
	var tests []hasherTest
	hs256 := hasherTest{Alg: "HS256", Header: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", expectedSign: "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
	tests = append(tests, hs256)
	for _, e := range tests {
		h := GetHasher(e.Alg)
		body := e.Header + "." + payload
		sign := h.Sign(body, secret)
		escapedSign := utils.EscapeNonUrlChars(sign)
		if escapedSign != e.expectedSign {
			t.Errorf("%s Hasher returned wrong signature. Expected: %s, got %s", e.Alg, e.expectedSign, escapedSign)
		}
	}
}

type hasherTest struct {
	Alg          string
	Header       string
	expectedSign string
}
