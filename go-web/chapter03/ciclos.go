// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nombre   string
	Duracion int
}

type Usuario struct {
	UserName      string
	Edad          int
	Activo        bool
	Administrador bool
	Cursos        []Curso
}

func FnTemplateHTML(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/ciclos.html")
	if err != nil {
		panic(err)
	}

	cursos := []Curso{Curso{"Python", 2}, Curso{"Go", 3}, Curso{"Angular", 5}}
	usuario := Usuario{
		UserName:      "Ernesto",
		Edad:          24,
		Activo:        true,
		Administrador: false,
		Cursos:        cursos}

	template.Execute(w, usuario)
}

func main() {
	http.HandleFunc("/", FnTemplateHTML)

	http.ListenAndServe("localhost:3000", nil)
}
