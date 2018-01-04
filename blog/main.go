package main

import (
	"fmt"
	"net/http"

	"github.com/DiegoSantosWS/blog/manipulador"
)

func main() {
	http.HandleFunc("/style.css", manipulador.CssPrincipal)
	http.HandleFunc("/", manipulador.Home)
	http.HandleFunc("/funcao", manipulador.Funcao)
	fmt.Println("Escutando...")
	http.ListenAndServe(":8181", nil)
}
