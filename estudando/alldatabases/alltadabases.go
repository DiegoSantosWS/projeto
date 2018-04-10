package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	//mysqldump -u root -p --all-databases > alldb_backup.sql
	//executa bk de todos os bancos em um unico arquivo
	cmd := exec.Command("mysqldump", "-uroot", "-p1234", "--all-databases")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Erro ao criar o arquivo")
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		log.Fatal()
	}
	err = ioutil.WriteFile("../bk/alldatabases.sql", bytes, 0644)
	if err != nil {
		panic(err)
	}
}
