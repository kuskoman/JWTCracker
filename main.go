package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kuskoman/JWTCracker/alghoritms"
	"github.com/kuskoman/JWTCracker/generators"
	"github.com/kuskoman/JWTCracker/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide a token to be cracked")
	}

	token := os.Args[1]
	alg := alghoritms.RecogniseJWTAlghoritm(token)
	if alg != "HS256" {
		log.Fatal("This alghoritm is not supported yet")
	}
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var chs []int
	chs = append(chs, 0)
	w := generators.WordGenerator{Alphabet: alphabet, CurrChs: chs}
	segments := strings.Split(token, ".")
	body := strings.Join(segments[0:2], ".")
	signature := utils.EscapeNonUrlChars(segments[2])
	i := 0
	for {
		currSig := w.Next()
		log.Printf("Current signature: %s\n", currSig)
		currHash := signHS256(body, currSig)
		currHash = utils.EscapeNonUrlChars(currHash)

		if currHash == signature {
			fmt.Printf("Signature found on %d iteration: %s", i, currSig)
			break
		}
		i++
	}
}

func signHS256(body, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	hash := mac.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash)
}
