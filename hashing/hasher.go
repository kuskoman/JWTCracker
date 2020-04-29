package hashing

import (
	"log"

	"github.com/kuskoman/JWTCracker/alghoritms"
)

type Hasher interface {
	Sign(body, secret string) string
}

func GetHasher(alg string) Hasher {
	var h Hasher

	switch alg {
	case "HS256":
		h = &alghoritms.HS256Hasher{}
	case "HS384":
		h = &alghoritms.HS384Hasher{}
	default:
		log.Fatal("This alghoritm is not supported yet")
	}
	return h
}
