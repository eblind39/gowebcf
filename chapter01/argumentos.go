// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
)

func FnArgumentos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Otro", "Header")
	fmt.Fprintf(w, r.URL.RawQuery) // un string

	fmt.Fprintf(w, "  <>  ")
	// leer como un mapa
	nombre := r.URL.Query().Get("nombre")
	fmt.Fprintf(w, "El nombre es: "+nombre)

	// Eliminar parámetro recibido
	valores := r.URL.Query()
	valores.Del("apellido")
	// Agregar parámetro
	valores.Add("edad", "24")

	fmt.Fprintf(w, "  <>  ")
	apellido := r.URL.Query().Get("apellido")
	if apellido != "" {
		fmt.Fprintf(w, "Apellido: "+apellido)
	} else {
		fmt.Fprintf(w, "No existe el parámetro: apellido")
	}

	fmt.Fprintf(w, "  <>  ")
	// Leer encabezado - llamar con curl a como sigue:
	// curl localhost:3000/argumentos?nombre=tito&apellido=gutierrez -H "access_token:123"
	accessToken := r.Header.Get("access_token")
	fmt.Fprintf(w, "Header access_token: "+accessToken)

	// crear encabezado
	r.Header.Set("tiempo", "12000")

	//mapargs := r.URL.Query()
	//for clave, valor := range mapargs {
	//	fmt.Fprintf(w, "  <>  ")
	//	fmt.Fprintf(w, "Param: "+clave+" - Valor: "+valor)
	//}
}

func main() {
	http.HandleFunc("/argumentos", FnArgumentos)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
