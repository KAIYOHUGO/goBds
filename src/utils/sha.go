package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(v []byte) string {
	h := sha1.New()
	h.Write(v)
	return hex.EncodeToString(h.Sum(nil))
}
