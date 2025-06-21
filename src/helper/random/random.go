package random

import (
	"math/rand"
)

func RandomNumbers(min, max int) string {
	length := RandomIntInRange(min, max)
	digits := "0123456789"
	return RandomFromCharset(length, digits)
}

func RandomLetters(min, max int) string {
	length := RandomIntInRange(min, max)
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandomFromCharset(length, letters)
}

func RandomAlphanumerics(min, max int) string {
	length := RandomIntInRange(min, max)
	alnum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return RandomFromCharset(length, alnum)
}

func RandomIntInRange(min, max int) int {
	if max < min {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func RandomFromCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
