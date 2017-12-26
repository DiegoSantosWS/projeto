package main

import (
	"github.com/DiegoSantosWS/cursodego/SERVIDOR/manipulador"
	"fmt"
	"net/http"
)

func main()  {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintln(w, "Olá Mundo")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando comandante...")
	http.ListenAndServe(":8181", nil)
}