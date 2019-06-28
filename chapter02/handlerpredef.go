// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	redireccionar := http.RedirectHandler("/hola", http.StatusMovedPermanently)
	notfound := http.NotFoundHandler()

	http.Handle("/redirect", redireccionar)
	http.Handle("/notfound", notfound)
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde redireccionado.")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
