package main

import "fmt"
import "github.com/DiegoSantosWS/cursodego/FUNCOES_AVANCADO/matematica"

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2,2)

	fmt.Printf("Oresultado é: %d\n\n", resultado)
	resultado = matematica.Soma(3,3)
	fmt.Printf("O result da soma foi: %d\n\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12,3)
	fmt.Printf("O result da divisao foi: %d\n\n", resultado)

	resultado, resto := matematica.DivisorComResto(12,5)
	fmt.Printf("O result da divisao foi: %d e o resto foi: %d\r\n", resultado, resto)
}