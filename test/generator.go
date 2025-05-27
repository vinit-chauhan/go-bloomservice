package test

import "math/rand"

func GenerateStringsOfLength(length int, count int) []string {
	strings := make([]string, count)
	for i := 0; i < count; i++ {
		strings[i] = RandomString(length)
	}
	return strings
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
