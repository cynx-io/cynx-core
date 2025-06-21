package random

import (
	"math/rand"
)

func RandomNumbers(length int) string {
	digits := "0123456789"
	return RandomFromCharset(length, digits)
}

func RandomLetters(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandomFromCharset(length, letters)
}

func RandomAlphanumerics(length int) string {
	alnum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return RandomFromCharset(length, alnum)
}

func RandomFromCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
