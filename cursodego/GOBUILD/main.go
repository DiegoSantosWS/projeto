package main

import (
	"fmt"
	"github.com/DiegoSantosWS/cursodego/GOBUILD/model"
)

var cache map[string]model.Imovel

func main()  {

	cache = make(map[string]model.Imovel, 0)
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa
	
	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	
	ap := model.Imovel{}
	ap.Nome = "AP AZUL"
	ap.X = 18
	ap.Y = 25
	ap.SetValor(60000)

	cache[ap.Nome] = ap
	fmt.Println("Quantos itens tem no cache?", len(cache))


	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Quantos itens tem no cache?", len(cache))
}