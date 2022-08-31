package helpers

import (
	"crypto/sha256"
	"fmt"
)

func Hash(raw string) (string, error) {
	h := sha256.New()
	h.Write([]byte(raw))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil

}

func MatchHash(raw string, encodedHash string) (bool, error) {
	hash, err := Hash(raw)
	return hash == encodedHash, err
}
