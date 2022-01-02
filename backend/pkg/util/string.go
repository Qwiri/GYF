package util

import (
	"math/rand"
	"regexp"
	"strings"
)

///

//goland:noinspection ALL
const CharSet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789"

func GenerateRandomString(length int) string {
	var bob strings.Builder
	for i := 0; i < length; i++ {
		bob.WriteRune(rune(CharSet[rand.Intn(len(CharSet))]))
	}
	return bob.String()
}

///

var usernameExpr = regexp.MustCompile("^[A-Za-z0-9_\\-]{1,16}$")

func IsNameValid(username string) bool {
	return usernameExpr.MatchString(username)
}
