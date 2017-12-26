package main

import (
	"github.com/DiegoSantosWS/cursodego/JSON_UNMARSHAL/model"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)


func main()  {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para Servidor. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Servidor. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do Servidor. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}

}