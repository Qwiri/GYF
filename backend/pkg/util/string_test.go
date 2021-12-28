package util

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	for _, length := range []int{0, 10, 25} {
		t.Run(fmt.Sprintf("length_%d", length), func(t *testing.T) {
			rnd := GenerateRandomString(length)
			if len(rnd) != length {
				t.Errorf("GenerateRandomString() = %v, want %v", len(rnd), length)
			}
		})
	}
}

func TestIsNameValid(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{"empty", "", false},
		{"too long", "12345678901234567", false},
		{"16", "1234567890123456", true},
		{"spaces", "hello world", false},
		{"no spaces", "helloworld", true},
		{"special chars", "helloworld!", false},
		{"only special chars", "!§$%&", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNameValid(tt.username); got != tt.want {
				t.Errorf("IsNameValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
