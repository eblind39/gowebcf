// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnDefaultServerMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde DefaultServerMux.")
}
func FnDesdeNuevoMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde el Handler de nuevo mux creado.")
}

func main() {
	// si se dejara siempre la línea siguiente, estaríamos
	// indicando que se trabajará con el DefaultServerMux
	// de http al crear el ListenAndServe (argumento nil)
	http.HandleFunc("/defaultmux", FnDefaultServerMux)

	mux := http.NewServeMux()
	// se agrega un único Handler al mux creado
	mux.HandleFunc("/newmux", FnDesdeNuevoMux)

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
