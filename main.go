package main

import (
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
	h := alghoritms.GetHasher(alg)
	w := generators.NewAlphabeticNumerator()
	segments := strings.Split(token, ".")
	body := strings.Join(segments[0:2], ".")
	signature := utils.EscapeNonUrlChars(segments[2])
	i := 1
	for {
		currSig := w.Next()
		log.Printf("Current signature: %s\n", currSig)
		currHash := h.Sign(body, currSig)
		currHash = utils.EscapeNonUrlChars(currHash)

		if currHash == signature {
			fmt.Printf("Signature found on %d iteration: %s", i, currSig)
			break
		}
		i++
	}
}
