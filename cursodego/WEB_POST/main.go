package main

import (
	
	"fmt"
	"bytes"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/DiegoSantosWS/cursodego/WEB_POST/model"
)


func main()  {
	cliente := &http.Client{
		Timeout: time.Second*30,
	}
	
	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "DIEGO DOS SANTOS"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar json", err.Error())
		return
	}


	request, err := http.NewRequest("POST", "https://requestb.in/1940s521", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o request bin. Erro: ", err.Error())
		return
	}

	request.SetBasicAuth("fizz","buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o servico do request bin. Erro: ", err.Error())
		return
	}
	
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo retornado pelo request bin. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}

}