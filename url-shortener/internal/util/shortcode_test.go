package util

import (
	"strings"
	"testing"
)

func TestGenerateShortCode(t *testing.T) {

	fixedLetters := "k8s-"
	expectedLength := len(fixedLetters) + 5

	code := GenerateShortCode(fixedLetters)

	if len(code) != expectedLength {
		t.Errorf("Incorrect code length: expected %d, but got %d", expectedLength, len(code))
	}

	if !strings.HasPrefix(code, fixedLetters) {
		t.Errorf("Generated code does not start with the expected prefix: expected '%s', got '%s'", fixedLetters, code)
	}

}
