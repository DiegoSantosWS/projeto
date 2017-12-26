package matematica

/*
Calculo executa qualquer tipo de calculo basta enviar a func desejada
*/
func Calculo(funcao func(int, int)int, numeroA int, nomeroB int) int {
	return funcao(numeroA, nomeroB)
}
//Efetuando multiplicacao
func Multiplicador(x int, y int) int {
	return x * y
}
//Efetua a Divisão
func Divisor(numeroA int, nomeroB int) (resultado int) {
	resultado = numeroA / nomeroB
	return
}
//Divisor com resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}