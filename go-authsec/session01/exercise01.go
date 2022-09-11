package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", encode)
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
