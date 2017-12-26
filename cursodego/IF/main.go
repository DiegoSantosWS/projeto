package main

import "fmt"

func main()  {


	meses := 0
	situacao := true
	cidade := "teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo")
	} 

	if situacao {
		fmt.Println("Esse credor esta inadimplete")
	}

	if !situacao {
		fmt.Println("Esse credor esta adimplente")
	}

	if cidade == "teste" {
		fmt.Println("O cliente vive no estado Teste")
	}

	if descricao, status := haQuantotempoEstaDevendo(meses); status {
		fmt.Println("Qual a situacao do clientes?", descricao)
	}

	fmt.Println("Obrigado por nos consultar")
}


func haQuantotempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "Cliente está devendo"
		return
	}
	descricao = "O cliente esta com o carner em dia"
	return
}