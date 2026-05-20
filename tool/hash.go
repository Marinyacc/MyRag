package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}
