package alghoritms

import "testing"

func TestAlghoritmRecognision(t *testing.T) {
	var exps []Expectation
	hs384Exp := Expectation{Alg: "HS384", Token: "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.RGFdh_VuEuURSubru7xP4rbaA4boUyueI7rEm75l1cNdE9gQ7H6mx2DYpauBjX5S"}
	hs256Exp := Expectation{Alg: "HS256", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
	exps = append(exps, hs384Exp, hs256Exp)

	for _, e := range exps {
		recognised := RecogniseJWTAlghoritm(e.Token)
		if recognised != e.Alg {
			t.Errorf("Wrong alghoritm recognised. Expected %s, got %s", e.Alg, recognised)
		}
	}
}

type Expectation struct {
	Token string
	Alg   string
}
