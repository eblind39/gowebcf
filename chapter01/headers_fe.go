// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnDeclarada(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("token", "123456")
	fmt.Fprintf(w, "Prueba de retorno de headers del back-end.")
}

func main() {
	http.HandleFunc("/headersbe", FnDeclarada)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
