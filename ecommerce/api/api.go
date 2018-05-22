package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DiegoSantosWS/ecommerce/models"
)

func MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
}

func Products(w http.ResponseWriter, r *http.Request) {
	var produto1 models.Produtos
	var prod []models.Produtos

	produto1 = models.Produtos{
		ID:          1,
		Name:        "PRODUTOS 1",
		Description: "PRODUTOS TESTE",
		Details:     "TESTE TESTE TESTE TESTE",
		Stock:       10.000,
		Type:        "UN",
		Sizer: models.Sizer{
			Tamanho: []string{"p", "m", "g"},
		},
		Price: models.Price{
			Promotion: 12.44,
			Price:     10.00,
		},
	}
	produto2 := models.Produtos{
		ID:          2,
		Name:        "PRODUTOS 2",
		Description: "PRODUTOS TESTE",
		Details:     "TESTE TESTE TESTE TESTE",
		Stock:       20.000,
		Type:        "UN",
		Sizer: models.Sizer{
			Tamanho: []string{"p", "m", "g"},
		},
		Price: models.Price{
			Promotion: 10.44,
			Price:     13.00,
		},
	}
	produto3 := models.Produtos{
		ID:          3,
		Name:        "PRODUTOS 3",
		Description: "PRODUTOS TESTE",
		Details:     "TESTE TESTE TESTE TESTE",
		Stock:       30.000,
		Type:        "UN",
		Sizer: models.Sizer{
			Tamanho: []string{"34", "35", "40"},
		},
		Price: models.Price{
			Promotion: 200.99,
			Price:     550.58,
		},
	}
	prod = append(prod, produto1, produto2, produto3)
	pJSON, err := json.Marshal(prod)
	if err != nil {
		log.Fatal("ERRO API: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(pJSON)
	fmt.Println(string(pJSON))
}
