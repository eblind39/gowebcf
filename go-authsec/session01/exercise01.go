package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":3007", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Name: "Ernst",
	}

	p2 := person{
		Name: "Alvison",
	}

	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("Bad data to encode", err)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	var people []person

	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Bad data to decode", err)
	}

	fmt.Println(people)
}