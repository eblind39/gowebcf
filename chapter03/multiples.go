// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type UsuarioF struct {
	UserName      string
	Edad          int
	Activo        bool
	Administrador bool
}

func FnTemplateHTMLMult(w http.ResponseWriter, r *http.Request) {
	funciones := template.FuncMap{
		"hola": hola,
		"suma": suma,
	}

	template, err := template.New("multiples.html").Funcs(funciones).ParseFiles("templates/multiples.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}

	usuario := UsuarioF{
		UserName:      "Ernesto",
		Edad:          24,
		Activo:        true,
		Administrador: true}

	template.Execute(w, usuario)
}

func (this UsuarioF) TienePermisoAdministrador(llave string) bool {
	return this.Activo && this.Administrador && llave == "ok"
}

func hola() string {
	return "Hola desde una función."
}

func suma(v1 int, v2 int) int {
	return v1 + v2
}

func main() {
	http.HandleFunc("/", FnTemplateHTMLMult)

	http.ListenAndServe("localhost:3000", nil)
}
