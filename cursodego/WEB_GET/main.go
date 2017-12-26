package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)


func main()  {
	cliente := &http.Client{
		Timeout: time.Second*30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Google Brasil", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler conteudo da pagina do Google Brasil", err.Error())
			return
		}

		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar request para o Google Brasil", err.Error())
		return
	}

	request.SetBasicAuth("teste","teste")
	resposta, err = cliente.Do(request)
	
	defer resposta.Body.Close()
	
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do Google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}