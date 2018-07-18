package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Exemple1Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá domain api1.dominio.com.br")
}

func Exemple2Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá domain api1.dominio.com.br")
}

func main() {
	r := mux.NewRouter()

	s1 := r.Host("api1.dominio.com.br").Subrouter()
	s2 := r.Host("api2.dominio.com.br").Subrouter()
	s1.HandleFunc("/teste", Exemple1Index)
	s2.HandleFunc("/", Exemple2Index)

	http.ListenAndServe(":3000", r)
}
