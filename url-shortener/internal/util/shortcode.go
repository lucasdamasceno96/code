package util

import (
	"crypto/rand"
	"fmt"
)

func randomDigits(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	s := ""
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%d", int(b[i]%10))
	}
	return s
}

func GenerateShortCode(fixedLetters string) string {
	digits := randomDigits(5)
	return fixedLetters + digits
}
