// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnLecturaMetodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Otro", "Header")
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET")
	case "POST":
		fmt.Fprintf(w, "POST")
	case "PUT":
		fmt.Fprintf(w, "PUT")
	case "DELETE":
		fmt.Fprintf(w, "DELETE")
	default:
		http.Error(w, "El método no existe.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/metodos", FnLecturaMetodos)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
