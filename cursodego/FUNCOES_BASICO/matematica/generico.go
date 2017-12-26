package matematica

/*
Calculo executa qualquer tipo de calculo basta enviar a func desejada
*/
func Calculo(funcao func(int, int)int, numeroA int, nomeroB int) int {
	return funcao(numeroA, nomeroB)
}

func Multiplicador(x int, y int) int {
	return x * y
}