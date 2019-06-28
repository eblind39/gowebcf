// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

func FnTemplateHTML(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/variables.html")
	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", FnTemplateHTML)

	http.ListenAndServe("localhost:3000", nil)
}
