package main

import (
	"math/rand"
	"strings"
)

const CharSet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789"

func GenerateRandomString(length int) string {
	var bob strings.Builder
	for i := 0; i < length; i++ {
		bob.WriteRune(rune(CharSet[rand.Intn(len(CharSet))]))
	}
	return bob.String()
}
