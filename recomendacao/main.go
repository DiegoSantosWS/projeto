package main

import (
	"fmt"

	con "github.com/DiegoSantosWS/recomendacao/connection"
	f "github.com/DiegoSantosWS/recomendacao/functions"
)

func init() {
	err := con.Connection()
	if err != nil {
		fmt.Println("Erro ao conectar com banco de dados: ", err.Error())
	}
	fmt.Println("Seja bem vindo")

}

var (
	idPedido  int
	idCliente int
	idProduto int
)

var param int = 30

func main() {
	//	fmt.Println("Sitema rodando")
	sql := "SELECT p.id, p.idCliente, pp.idProduto FROM sig_pedidos AS p INNER JOIN sig_pedidos_produtos AS pp ON p.id=pp.idPedido WHERE p.status = 'Aprovado' ORDER BY p.id DESC"
	rows, err := con.Db.Queryx(sql)
	if err != nil {
		fmt.Println("ERROR[MAIN]: pedido não encontrado", err.Error())
		return
	}
	for rows.Next() {
		err = rows.Scan(&idPedido, &idCliente, &idProduto)
		if err != nil {
			fmt.Println("ERROR[MAIN]: Error ao realizar scan do pedido. ", err.Error())
			return
		}
		//Retornando nome do cliente
		name, err := f.GetNameCliente(idCliente)
		if err != nil {
			fmt.Println("ERROR[MAIN]: Nome do cliente não encontrado", err.Error())
		}
		//Retornando email do cliente
		email, err := f.GetEmailCliente(idCliente)
		if err != nil {
			fmt.Println("ERROR[MAIN]: E-mail do cliente não encontrado", err.Error())
		}
		//Retornando nome do produto
		nProduto, err := f.GetNomeProdutos(idProduto)
		if err != nil {
			fmt.Println("ERROR[MAIN]: Nome do produto não encontrado", err.Error())
		}
		//Retornando quantidade de dias da ultima compra realizada pelo cliente
		uCompra, _ := f.GetDiasUltimaCompra(idCliente)
		//Retorna a categoria dos produtos relacionado a ultima compra do cliente
		cate, _ := f.GetCategoriaProduto2(idProduto)
		/**
		 * Analisando se a ultima compra do cleinte tem  mais de 30 dias,
		 * Se o retorno for maior sistema vai disparar um email para o cliente
		 * caso contrario ele dsbot não fará nada.
		 */
		if uCompra > param {
			//Inicia o envio do email com a recomendação para o cliente
			f.EnviaEmail(name, email, nProduto, uCompra, cate, idProduto)
		}
	}
}
