package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var key = []byte{}

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

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

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)

	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}
