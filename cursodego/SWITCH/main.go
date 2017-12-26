package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	
	numero := 3
	fmt.Print("O numero ", numero, " Se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	}

	fmt.Println("Você é da familha do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim")
	default:
		fmt.Println("Deixa Pra la")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Você pode ficar dormindo até mais tarde")
	default:
		fmt.Println("Vai trabalhar Vagabundo")
	}

	numero = 10
	fmt.Println("Este numero cabe no Digito?")
	switch {
	case numero > 10:
		fmt.Println("Sim")
	case numero>=10 && numero <100:
		fmt.Println("Serve dois dígitos...")
	case numero <=100:
		fmt.Println("Não")
		
	}
}