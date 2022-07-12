// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"log"
	"net/http"
)

func FnEnvioErrores(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Otro", "Header")
	http.Error(w, "Este es un error", http.StatusConflict)
}

func main() {
	http.HandleFunc("/errores", FnEnvioErrores)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
