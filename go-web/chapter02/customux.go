// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

// se crea el tipo Handler
type customHandler func(http.ResponseWriter, *http.Request)

// el mux que encapsula las rutas con sus handlers asociados
// en el formato clave/valor dentro de un mapa
// clave: ruta
// valor: Handler
type CustomMux struct {
	rutas map[string]customHandler // handlers
}

// implementar la interface ServeHTTP, para que cada función
// dentro del mapa sea interpretado como un Handler
func (this *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path) // clave del mapa

	if fn, ok := this.rutas[r.URL.Path]; ok {
		fn(w, r)
	} else {
		fmt.Println("URL no encontrada.")
		http.NotFound(w, r)
	}
}

// se agrega un nuevo handler al mapa de CustomMux
func (this *CustomMux) AddMux(ruta string, fn customHandler) {
	this.rutas[ruta] = fn
}

func main() {
	newMap := make(map[string]customHandler)
	mux := &CustomMux{rutas: newMap}

	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde una función anónima")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
