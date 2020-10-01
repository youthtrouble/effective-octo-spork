package crypto 

import (
	"crypto/sha256"
)

func newHash256(data []byte) []byte {
	newhash := sha256.Sum256(data)
	return newhash[:]
}

func newHash224(data []byte) []byte {
	newHash := sha256.Sum224(data)
	return newHash[:]
}