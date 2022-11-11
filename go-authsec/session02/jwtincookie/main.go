package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	"time"
)

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey = "i love thursdays when it rains 8723 inches"

func main() {
	http.HandleFunc("/", everything)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}

func everything(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	ss := c.Value
	afterVerificationToken, err := jwt.ParseWithClaims(ss, &myClaims{}, func(beforeVerificationToken *jwt.Token) (interface{}, error) {
		if beforeVerificationToken.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("SOMEONE TRIED TO HACK changed signing method")
		}

		return []byte(myKey), nil
	})

	// StandardCLaims has the...
	// Valid() error
	// ... method wich means it implements the Claims interface...
	/*
		type Claims interface {
			Valid() error
		}
	*/
	// ...when you ParseWithClaims...
	// the Valid() method gets run
	// ... and if all is well, then returns no "error" and
	// type TOKEN which has a field VALID will be true

	isEqual := err == nil && afterVerificationToken.Valid

	message := "Not logged in"
	if isEqual {
		message = "Logged in"
		claims := afterVerificationToken.Claims.(*myClaims)
		fmt.Println(claims.Email)
		fmt.Println(claims.ExpiresAt)
	}

	html := `<!DOCTYPE html>
			<html lang="en">
			  <head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>HMAC Example</title>
			  </head>
			  <body>
				<p>Cookie value: ` + c.Value + `</p>
				<p>` + message + `</p>
				<form action="/submit" method="post">
					<input type="text" name="emailjwt" />
					<input type="submit" />
				</form>
			  </body>
			</html>`

	io.WriteString(w, html)
}

func getJWT(msg string) (string, error) {
	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("couldn't SignedString in %w", err)
	}

	return ss, nil
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailjwt")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "couldn't getJWT", http.StatusInternalServerError)
		return
	}

	// "hash / message digest / digest / hash value" | "what we stored"
	c := http.Cookie{
		Name:  "session",
		Value: ss,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
