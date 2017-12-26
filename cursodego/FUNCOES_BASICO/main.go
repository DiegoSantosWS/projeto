package main

import "fmt"
import "github.com/DiegoSantosWS/cursodego/FUNCOES_BASICO/matematica"

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2,2)

	fmt.Printf("Oresultado é: %d\n\n", resultado)
	resultado = matematica.Soma(3,3)
	fmt.Printf("O result da soma foi: %d\n\n", resultado)
}

