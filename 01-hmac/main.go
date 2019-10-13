package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	hashed := "example@mail.com"
	fmt.Println(getHash(hashed))

	hashed = "example@mail.com"
	fmt.Println(getHash(hashed))
}

func getHash(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
