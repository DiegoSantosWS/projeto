package main

import (
	"sync"
	"time"
)

func main() {
	/*
	   Exemplo de goroutines controladas através de waitGroup
	   waitGroup aguarda a execução de goroutines até que elas sejam finalizadas
	*/

	println("Inicio")
	//Cria um grupo de espera para aguardar o processamento de todas as goroutines
	var waitGroup sync.WaitGroup

	//Define o numero de goroutines que deve ser aguardado.
	waitGroup.Add(3) /*Mude para 3 para aguardar as 3 goroutines ou mode para 0 para não aguardar nenhuma delas*/
	go func() {
		time.Sleep(2 * time.Second)
		println("-> foo")
		//Decrementando o contador
		waitGroup.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		println("-> bar")
		//Decrementando o contador
		waitGroup.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		println("-> qux")
		//Decrementando o contador
		waitGroup.Done()
	}()
	//Segura a execução até que o contador WaitGroup chegue a 0
	waitGroup.Wait()
	println("Fim")
}
