package util

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
