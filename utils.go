package main

import (
	"math/rand"
	"unicode/utf8"
)

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomRuneFromString(str string) string {
	n := utf8.RuneCountInString(str)
	r := []rune(str)[rand.Intn(n)]
	return string(r)
}
