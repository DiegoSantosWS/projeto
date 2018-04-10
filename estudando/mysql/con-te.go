package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db recebe um ponteiro de sqlx.DB
var Db *sqlx.DB

//Conection abre uma conexão com banco de dados
func Conection() (err error) {
	err = nil

	Db, err = sqlx.Open("mysql", "shortnerurl:synataseb@tcp(209.126.122.28:3306)/zadmin_shortnerurl")
	if err != nil {
		fmt.Println("Erro ao conectar", err.Error())
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar", err.Error())
		return
	}
	return
}

func main() {
	err := Conection()
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}

	_, err = Db.Exec("DROP TABLE IF EXISTS `logquery`;")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	_, err = Db.Exec("CREATE TABLE `logquery` (`id` int(11) NOT NULL AUTO_INCREMENT,`url` varchar(255) DEFAULT NULL,`token` varchar(45) DEFAULT NULL,`ip` varchar(45) DEFAULT NULL,`data` datetime DEFAULT NULL,`referencia` varchar(200) DEFAULT NULL,`browser` varchar(2000) DEFAULT NULL,`sysoperacional` varchar(50) DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	_, err = Db.Exec("DROP TABLE IF EXISTS `url`;")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	_, err = Db.Exec("CREATE TABLE `url` (`id` int(11) NOT NULL AUTO_INCREMENT,`url` varchar(255) NOT NULL,`tokenMD5` varchar(32) NOT NULL,`token` varchar(255) NOT NULL,`shortenURL` varchar(45) NOT NULL,PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	_, err = Db.Exec("DROP TABLE IF EXISTS `users`")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	_, err = Db.Exec("CREATE TABLE `users` (`id` int(11) NOT NULL AUTO_INCREMENT,`nome` varchar(45) NOT NULL,`email` varchar(45) NOT NULL,`login` varchar(45) NOT NULL,`pass` varchar(255) NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
		return
	}
	fmt.Println("Conectado criado")

}
