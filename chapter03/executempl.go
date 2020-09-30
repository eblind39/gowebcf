// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"html/template"
	"net/http"
)

type UsuarioEx struct {
	UserName      string
	Edad          int
	Activo        bool
	Administrador bool
}

// Cargar todos los templates en el server y utilizarlos luego bajo demanda
var templates = template.Must(template.New("T").ParseFiles("templates/executempl.html", "templates/footer.html"))

// Cuando ya se trabaja con demasiados templates, se debe dejar de usar
// ParseFiles y usar ParseGlob, lo que hace es cargar todos los templates
// de un directorio, sin embargo, los templates que se necesiten cargar
// corresponden a los utilizados dentro del programa Go en específico.
//   var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))

// acceder a los subdirectorios.
//   var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

func FnTemplateHTMLExec(w http.ResponseWriter, r *http.Request) {
	usuario := UsuarioEx{
		UserName:      "Ernesto",
		Edad:          24,
		Activo:        true,
		Administrador: true,
	}

	err := templates.ExecuteTemplate(w, "executempl", usuario)
	if err != nil {
		panic(err)
	}
}

func (this UsuarioEx) TienePermisoAdministrador(llave string) bool {
	return this.Activo && this.Administrador && llave == "ok"
}

func main() {
	http.HandleFunc("/", FnTemplateHTMLExec)

	http.ListenAndServe("localhost:3000", nil)
}
