package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func main() {
	res := ValidMAC([]byte("00-08-20-83-53-D1"), []byte("00-08-20-83-53-D1"), []byte("123456"))
	if res {
		println("true")
	}
	println("false")
}
