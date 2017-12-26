package main

import (
	"fmt"
	"github.com/DiegoSantosWS/cursodego/pacotes/operadora"
	"github.com/DiegoSantosWS/cursodego/pacotes/prefixo"
)
//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "jeff"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}