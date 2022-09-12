package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	// curl -u "username:password" -v google.com
	// Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("username:password")))

	pw := "123456789"
	hashedPw, err := hashPassword(pw)
	if err != nil {
		panic(err)
	}

	comparePassword(pw, hashedPw)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating brcypt hash from password: %q", password)
	}

	return bs, nil
}

func comparePassword(password string, hashedPW []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), hashedPW)
	if err != nil {
		return fmt.Errorf("invalid password: %q", err)
	}

	return nil
}
