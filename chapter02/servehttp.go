// programa de ejemplo donde se demuestra el uso de http
// implementación de la interface ServeHTTP de un mux
package main

import (
	"fmt"
	"net/http"
)

type Usuario struct {
	nombre string
}

func (this *Usuario) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola: "+this.nombre)
}

func main() {
	ernesto := &Usuario{nombre: "Ernesto"}
	servidor := &http.Server{Addr: "localhost:3000", Handler: ernesto}

	// averiguar cómo utilizarlo al iniciar el servidor
	//log.Fatal(http.ListenAndServe("localhost:3000", servidor))
}
