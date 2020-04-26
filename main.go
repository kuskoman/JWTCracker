package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a token to be cracked")
	}
	token := os.Args[1]
	println(token)
}

func parseToken(token string) string {
	h := &Header{}
	segments := strings.Split(token, ".")

	if len(segments) != 3 {
		log.Fatal("Token has less or more than 3 segments. Make sure you are providing a signed JWS")
	}

	decodedHeaderJson, err := base64.StdEncoding.DecodeString(segments[0])
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
