// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnHolaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo! - redirigido.")
}
func FnDeclarada(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/hola", http.StatusMovedPermanently)
	//if err != nil {
	//	htttp.NotFound(w, r)
	//}
}

func main() {
	http.HandleFunc("/redirect", FnDeclarada)
	http.HandleFunc("/hola", FnHolaMundo)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
