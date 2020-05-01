package hashing

import (
	"crypto/sha256"
	"crypto/sha512"
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
		h = &alghoritms.SHAHasher{Alg: sha256.New}
	case "HS384":
		h = &alghoritms.SHAHasher{Alg: sha512.New384}
	default:
		log.Fatal("This alghoritm is not supported yet")
	}
	return h
}
