// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnDeclarada(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Llamada desde una Fn declarada.")
}

func main() {

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo!")
	})

	http.HandleFunc("/alterno", FnDeclarada)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
