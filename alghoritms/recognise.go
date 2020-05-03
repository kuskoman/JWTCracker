package alghoritms

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
)

func RecogniseJWTAlghoritm(token string) string {
	h := &Header{}
	segments := strings.Split(token, ".")

	if len(segments) != 3 {
		log.Fatal("Token has less or more than 3 segments. Make sure you are providing a signed JWS")
	}

	decodedHeaderJson, err := base64.RawStdEncoding.DecodeString(segments[0])
	if err != nil {
		log.Fatal("Error when decoding token header")
	}

	err = json.Unmarshal(decodedHeaderJson, h)
	if err != nil {
		log.Fatal("Error when parsing header JSON")
	}

	return h.Algorithm
}

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}
