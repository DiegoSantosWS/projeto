package functions

import (
	"fmt"

	con "github.com/DiegoSantosWS/recomendacao/connection"
)

//GetNameCliente Recebe Codigo e retorna nome do cliente
func GetNameCliente(codclient int) (string, error) {
	var name string
	sql := "SELECT nome FROM sig_clientes WHERE id = ? "
	row, err := con.Db.Queryx(sql, codclient)
	if err != nil {
		fmt.Println("ERROR: nome do cliente não encontrado", err.Error())
		return "", err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			fmt.Println("ERROR: nome do cliente não encontrado", err.Error())
			return "", err
		}
	}
	return string(name), nil
}

//GetEmailCliente recebe o codigo do cliente RETORNA EMAIL DO CLIENTE
func GetEmailCliente(codClient int) (string, error) {
	var email string
	sql := "SELECT email FROM sig_clientes_emails WHERE principal = '1' AND idCliente = ? ORDER BY id DESC LIMIT 1"
	row, err := con.Db.Queryx(sql, codClient)
	if err != nil {
		fmt.Println("ERROR: email do cliente não encontrado", err.Error())
		return "", err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&email)
		if err != nil {
			fmt.Println("ERROR: email do cliente não encontrado", err.Error())
			return "", err
		}
	}
	return string(email), nil
}

//GetNomeProdutos recebe codigo RETORNA NOME DO PRODUTO
func GetNomeProdutos(idproduto int) (string, error) {
	var nameProduct string
	sql := "SELECT nome FROM sig_produtos WHERE id = ?"
	row, err := con.Db.Queryx(sql, idproduto)
	if err != nil {
		fmt.Println("ERROR: Produto não encontrado", err.Error())
		return "", err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&nameProduct)
		if err != nil {
			fmt.Println("ERROR: Produto não encontrado não encontrado", err.Error())
			return "", err
		}
	}
	return string(nameProduct), nil
}

//GetCategoriaProduto recebe o codigo do produto e retorna a categoria cadastrada nele
func GetCategoriaProduto(idproduto int) (string, error) {
	var idCategoria int
	sql := "SELECT categoria FROM sig_produtos WHERE id = ?"
	rows, err := con.Db.Queryx(sql, idproduto)
	if err != nil {
		fmt.Println("Erro: categoria do produto não foi encontrada", err.Error())
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&idCategoria)
		if err != nil {
			fmt.Println("ERROR: categoria do produto não foi encontrada ", err.Error())
			return "", err
		}
	}
	nameCategoria, _ := GetNameCategoria(idCategoria)
	return string(nameCategoria), nil
}

//GetNameCategoria recebe o id e retornar nome da categoria
func GetNameCategoria(id int) (string, error) {
	var nameCategoria string
	sql := "SELECT nome FROM sig_categorias WHERE id = ?"
	rows, err := con.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("Erro: categoria do produto não foi encontrada", err.Error())
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&nameCategoria)
		if err != nil {
			fmt.Println("ERROR: categoria do produto não foi encontrada ", err.Error())
			return "", err
		}
	}
	return string(nameCategoria), nil
}

//GetProdutosPorCategorias recebe o id da categoria e retorna todos os produtos cadastrados nela
//GetProdutosPorCategorias recebe o id da categoria e retorna todos os produtos cadastrados nela
func GetProdutosPorCategorias(id int) (interface{}, error) {
	var (
		idP  int
		name string
		slug string
	)
	categoria, err := GetCategoriaProduto2(id)
	if err != nil {
		fmt.Println("ERROR: Produto não encontrado", err.Error())
		return "", err
	}
	sql := "SELECT id, nome, slug FROM sig_produtos WHERE categoria = ? AND id != ? LIMIT 10"
	row, err := con.Db.Queryx(sql, categoria, id)
	if err != nil {
		fmt.Println("ERROR: Produto não encontrado", err.Error())
		return "", err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&idP, &name, &slug)
		if err != nil {
			fmt.Println("ERROR: Produto não encontrado não encontrado", err.Error())
			return "", err
		}
	}
	data := map[string]interface{}{
		"id":   idP,
		"Nome": name,
		"slug": slug,
	}
	return data, nil
}

//GetCategoriaProduto2  recebe o id da categoria e retorna todos os produtos cadastrados nela
func GetCategoriaProduto2(idproduto int) (int, error) {
	var idCategoria int
	sql := "SELECT categoria FROM sig_produtos WHERE id = ?"
	rows, err := con.Db.Queryx(sql, idproduto)
	if err != nil {
		fmt.Println("Erro: categoria do produto não foi encontrada", err.Error())
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&idCategoria)
		if err != nil {
			fmt.Println("ERROR: categoria do produto não foi encontrada ", err.Error())
			return 0, err
		}
	}
	return int(idCategoria), nil
}

//GetDiasUltimaCompra recebe codigo do cliente para contar quantos dias tem entre a data atual
//e a ultima compra do cliente
func GetDiasUltimaCompra(codCliente int) (int, error) {
	var dias int
	sql := "SELECT DATEDIFF(DATE_FORMAT(CURDATE(), '%Y%m%d'), DATE_FORMAT(dataPedido, '%Y%m%d')) as ultima FROM sig_pedidos AS p INNER JOIN sig_pedidos_produtos AS pp ON p.id=pp.idPedido WHERE p.status = 'Aprovado' AND idCliente = ? ORDER BY p.id DESC"
	rows, err := con.Db.Queryx(sql, codCliente)
	if err != nil {
		fmt.Println("ERROR: dias", err.Error())
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dias)
		if err != nil {
			fmt.Println("ERROR: dias", err.Error())
			return 0, err
		}
	}
	return int(dias), nil

}

//NameSlugCategoria para algumas lojas essa função será muito util pois retorna o slug do produto
func NameSlugCategoria(id int) (string, error) {
	var nameCategoria string
	sql := "SELECT slug FROM sig_categorias WHERE id = ?"
	rows, err := con.Db.Queryx(sql, id)
	if err != nil {
		fmt.Println("Erro: categoria do produto não foi encontrada", err.Error())
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&nameCategoria)
		if err != nil {
			fmt.Println("ERROR: categoria do produto não foi encontrada ", err.Error())
			return "", err
		}
	}
	return string(nameCategoria), nil
}
