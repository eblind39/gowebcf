package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	msg := "This is totally fun get hands-on and learning it from the ground up."
	password := "ilovedogs"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}
	bs = bs[:16]

	rslt, err := enDecode(bs, msg)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("before base64", string(rslt))

	rslt2, err := enDecode(bs, string(rslt))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rslt2))
}

func enDecode(key []byte, input string) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't newCipher %w", err)
	}

	// initialization vector
	iv := make([]byte, aes.BlockSize)
	s := cipher.NewCTR(b, iv)
	buff := &bytes.Buffer{}
	sw := cipher.StreamWriter{
		S: s,
		W: buff,
	}
	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't ws.Write to streamwriter %w", err)
	}

	return buff.Bytes(), nil
}
