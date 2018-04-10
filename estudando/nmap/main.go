package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("nmap", "-v", "-A", "127.0.0.1")
	maps, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Erro ao iniciar a execução")
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(maps)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		log.Fatal()
	}

	defer maps.Close()
	err = ioutil.WriteFile("../bk/map.sql", bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
