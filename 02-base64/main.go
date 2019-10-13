package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := `
	Beauty isn't seen by eyes.
	It's felt by hearts,
	Recognized by souls,
	In the presence of love.
	`

	bs64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(bs64)

	dbs64, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dbs64))
}
