package repo

import (
	"github.com/jmoiron/sqlx"
	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
