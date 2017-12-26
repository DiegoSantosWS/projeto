package main

import (
	"fmt"
	"os"
	"encoding/csv"
)


func main()  {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("Ouve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}
	/*
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("O conteudo da linha é:", linha)
	}
	*/

	leitor := csv.NewReader(arquivo)
	conteudo, err := leitor.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro:", err.Error())
		return
	}

	for indiceLinha, linha := range conteudo{
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indiceItem, item)	
		}
	}
	arquivo.Close()
}