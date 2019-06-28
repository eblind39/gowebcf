// programa de ejemplo donde se demuestra el uso de http
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func getURL() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic("Error irrecuperable")
	}
	u.Host = "localhost:3000"
	u.Scheme = "http" // puede ser https
	// agregar parámetros por query string
	query := u.Query()
	query.Add("nombre", "valor")
	u.RawQuery = query.Encode()
	return u.String()
}

func FnGeneraURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getURL())
}

func main() {
	http.HandleFunc("/generaurl", FnGeneraURL)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
