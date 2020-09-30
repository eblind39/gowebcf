// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type UsuarioRndr struct {
	UserName string
}

type ErrorC struct {
	ErrorTitleC string
	ErrorCodeC  int
	ErrorMsgC   string
}

var templates = template.Must(template.New("T").ParseFiles("templates/rendertempl.html", "templates/footer.html"))
var errortempl = template.Must(template.ParseFiles("templates/error.html"))

func handleError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	errorc := ErrorC{
		ErrorTitleC: "Error en rendertempl.go",
		ErrorCodeC:  status,
		ErrorMsgC:   msg,
	}
	errortempl.Execute(w, errorc)
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "se debe pasar err convertido en string")
		//http.Error(w, "No es posible retornar el template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "usuario", nil)
	})

	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		usuario := UsuarioRndr{UserName: "Ernesto"}
		renderTemplate(w, "usuario", usuario)
	})

	http.ListenAndServe("localhost:3000", nil)
}
