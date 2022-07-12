// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type Usuario struct {
	UserName string
	Edad     int
}

func FnTemplateHTML(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/templatehtml.html")
	if err != nil {
		panic(err)
	}

	usuario := Usuario{UserName: "Ernesto", Edad: 24}

	template.Execute(w, usuario)
}

func main() {
	http.HandleFunc("/", FnTemplateHTML)

	http.ListenAndServe("localhost:3000", nil)
}
