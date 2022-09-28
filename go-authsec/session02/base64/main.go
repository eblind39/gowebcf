package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "This is totally fun get hands-on and learning it from the ground up."
	encoded := encode(msg)
	fmt.Println("encoded msg: ", encoded)

	s, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("decoded msg: ", s)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}

	return string(s), nil
}
