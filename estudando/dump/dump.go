package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello world teste")
	st, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("erro ao executar comando", err.Error())
		log.Fatal(err)
	}

	outfile, err := os.Create("bk/Ola.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		log.Fatal(err)
	}

	defer outfile.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println("Erro ao iniciar a execução")
		log.Fatal(err)
	}

	io.Copy(outfile, st)
}
