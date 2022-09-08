package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type shape struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":3007", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	s1 := shape{
		Name: "Rectangle",
	}

	err := json.NewEncoder(w).Encode(s1)
	if err != nil {
		log.Println("Encoded bad data: ", err)
	}

}

func decode(w http.ResponseWriter, r *http.Request) {

}
