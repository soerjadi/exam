package utils

import (
	"crypto/rand"
	"fmt"
)

// RandString --
func RandString(length int) string {
	ran := make([]byte, length)
	rand.Read(ran)
	return fmt.Sprintf("%x", ran)
}
