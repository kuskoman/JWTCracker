package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a token to be cracked")
	}
	token := os.Args[1]
	//alg := getAlghoritm(token)

	segments := strings.Split(token, ".")
	body := strings.Join(segments[0:2], ".")
	signature := segments[2]
	a := sign(body, "a")
	fmt.Println(signature)
	fmt.Println(a)
}

func getAlghoritm(token string) string {
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

func sign(body, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}
