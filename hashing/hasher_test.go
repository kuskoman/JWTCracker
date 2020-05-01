package hashing

import (
	"testing"

	"github.com/kuskoman/JWTCracker/utils"
)

func TestHS256Hasher(t *testing.T) {
	h := GetHasher("HS256")
	payload := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
	secret := "your-256-bit-secret"
	sign := h.Sign(payload, secret)
	escapedSign := utils.EscapeNonUrlChars(sign)
	expectedSign := "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	if escapedSign != expectedSign {
		t.Errorf("HS256 Hasher returned wrong signature. Expected: %s, got %s", expectedSign, escapedSign)
	}
}
