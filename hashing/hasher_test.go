package hashing

import (
	"fmt"
	"testing"

	"github.com/kuskoman/JWTCracker/utils"
)

func TestHS256Hasher(t *testing.T) {
	payload := "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
	secret := "your-256-bit-secret"
	var tests []hasherTest
	hs256 := hasherTest{Alg: "HS256", Header: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", expectedSign: "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
	hs384 := hasherTest{Alg: "HS384", Header: "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9", expectedSign: "RGFdh_VuEuURSubru7xP4rbaA4boUyueI7rEm75l1cNdE9gQ7H6mx2DYpauBjX5S"}
	hs512 := hasherTest{Alg: "HS512", Header: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9", expectedSign: "pazba9Pj009HgANP4pTCQAHpXNU7pVbjIGff_plktSzsa9rXTGzFngaawzXGEO6Q0Hx5dtGi-dMDlIadV81o3Q"}
	tests = append(tests, hs256, hs384, hs512)
	for _, e := range tests {
		h := GetHasher(e.Alg)
		body := e.Header + "." + payload
		sign := h.Sign(body, secret)
		escapedSign := utils.EscapeNonUrlChars(sign)
		if escapedSign != e.expectedSign {
			fmt.Println(body)
			t.Errorf("%s Hasher returned wrong signature. Expected: %s, got %s", e.Alg, e.expectedSign, escapedSign)
		}
	}
}

type hasherTest struct {
	Alg          string
	Header       string
	expectedSign string
}
