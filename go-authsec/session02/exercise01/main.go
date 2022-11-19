package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
)

var db = map[string][]byte{}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("errormsg")

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Home</title>
			</head>
			<body>
				<h1>IF THERE WAS ANY ERROR, HERE IT IS: %s</h1>
				<form action="/register" method="POST">
					<input type="email" name="e">
					<input type="password" name="p">
					<input type="submit" >
				</form>
			</body>
		</html>
	`, errMsg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		erroMsg := url.QueryEscape("your method was not post")
		http.Redirect(w, r, "/?errormsg="+erroMsg, http.StatusSeeOther)
		return
	}

	e := r.FormValue("e")
	if e == "" {
		erroMsg := url.QueryEscape("your email needs to not be empty")
		http.Redirect(w, r, "/?errormsg="+erroMsg, http.StatusSeeOther)
		return
	}
	p := r.FormValue("p")
	if p == "" {
		erroMsg := url.QueryEscape("your password needs to not be empty")
		http.Redirect(w, r, "/?errormsg="+erroMsg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		erroMsg := "there was an internal server error - evil laugh: hahahah"
		http.Error(w, erroMsg, http.StatusInternalServerError)
		return
	}

	log.Println("password", p)
	log.Println("bcrypted", bsp)

	db[e] = bsp

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
