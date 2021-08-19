package naming

import (
	"math/rand"
	"unicode/utf8"
)

func RandomRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomRuneFromString(str string) string {
	n := utf8.RuneCountInString(str)
	r := []rune(str)[rand.Intn(n)]
	return string(r)
}

func contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}

	return false
}
