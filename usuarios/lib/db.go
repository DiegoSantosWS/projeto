package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var configuracao = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "1234",
	Database: "memboros",
}

//Sess é uma variavel que faz uma coneção com banco de dados
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(configuracao)
	if err != nil {
		log.Fatal(err.Error())
	}
}
