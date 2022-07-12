// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

func FnResponderConTemplate(w http.ResponseWriter, r *http.Request) {
	// no tiene funcionalidad importante porque no se invoca
	// un archivo de plantilla
	template, err := template.New("Hola").Parse("Hola Mundo!")
	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", FnResponderConTemplate)

	http.ListenAndServe("localhost:3000", nil)
}
