// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type UsuarioT struct {
	UserName      string
	Edad          int
	Activo        bool
	Administrador bool
}

func FnTemplateHTMLMethod(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/metodos.html")
	if err != nil {
		panic(err)
	}

	usuario := UsuarioT{
		UserName:      "Ernesto",
		Edad:          24,
		Activo:        true,
		Administrador: true}

	template.Execute(w, usuario)
}

func (this UsuarioT) TienePermisoAdministrador(llave string) bool {
	return this.Activo && this.Administrador && llave == "ok"
}

func main() {
	http.HandleFunc("/", FnTemplateHTMLMethod)

	http.ListenAndServe("localhost:3000", nil)
}
