package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("UNIX box?\r\n")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "plan9":
		fmt.Printf("YES!\r\n")
	case "linux":
		fmt.Printf("YES! I'm using operating system linux\r\n")
	default:
		fmt.Printf("Not at all\r\n")
	}

	fmt.Println("Checando número de 1 a 10\r\n")

	fmt.Print("Digite um número: ")
	var inserido string
	fmt.Scanln(&inserido)

	switch numero, _ := strconv.Atoi(inserido); numero {
	case 1, 3, 5, 7:
		fmt.Printf("%v é primo!\n\r", numero)
		fallthrough
	case 9:
		fmt.Printf("%v é impar!\n\r", numero)
		fmt.Printf("O resultado da divisão de %v por 2 é %v!\n\r", numero, numero%2)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%v é par!\n\r", numero)
	default:
		fmt.Printf("%v não está entre 1 e 10!\n\r", numero)
	}
	//fmt.Println(numero)
}
