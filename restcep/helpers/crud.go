package helpers

import (
	"encoding/json"
	"fmt"
	"log"

	con "github.com/DiegoSantosWS/restcep/connection"
	"github.com/DiegoSantosWS/restcep/model"
)

//InsertBD ...
func InsertBD(d *model.DadosCep) {
	sql := "INSERT INTO address "
	sql += "(cidade,bairro,logradouro,cep,estado,uf_area_km2,uf_codigo_ibge,uf_nome,cidade_area_km2, cidade_codigo_ibge) "
	sql += " VALUES (?,?,?,?,?,?,?,?,?,?) "
	ros, err := con.Db.Exec(sql, &d.Cidade, &d.Bairro, &d.Logradouro, &d.CEP, &d.Estado, &d.UFINFO.AreaKm2, &d.UFINFO.CodigoIbge, &d.UFINFO.Nome, &d.CidadeInfo.AreaKm2, &d.CidadeInfo.CodigoIbge)
	if err != nil {
		log.Fatal("ERRO ao inserir dados no banco de dados", err.Error())
		return
	}

	fmt.Println(ros.LastInsertId())
	return

}

//UpdateBD atualiza dados encontrado
func UpdateBD(d *model.DadosCep, cod string) bool {
	sql := "UPDATE address SET "
	sql += "cidade             = ?, "
	sql += "bairro 			   = ?, "
	sql += "logradouro		   = ?, "
	sql += "cep                = ?, "
	sql += "estado             = ?, "
	sql += "uf_area_km2        = ?, "
	sql += "uf_codigo_ibge	   = ?, "
	sql += "uf_nome 		   = ?, "
	sql += "cidade_area_km2    = ?, "
	sql += "cidade_codigo_ibge = ?  "
	sql += "WHERE cep = ?  "
	rows, err := con.Db.Exec(sql, d.Cidade, d.Bairro, d.Logradouro, d.CEP, d.Estado, d.UFINFO.AreaKm2, d.UFINFO.CodigoIbge, d.UFINFO.Nome, d.CidadeInfo.AreaKm2, d.CidadeInfo.CodigoIbge, cod)
	if err != nil {
		log.Fatal("ERRO ao alterar dados no banco de dados", err.Error())
		return false
	}

	_, err = rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO ao alterar dados no banco de dados", err.Error())
		return false
	}
	return true
}

//SelectBD busca os dados selecionados
func SelectBD(codigo string) (cepJSON interface{}) {
	sql := "SELECT * FROM address WHERE cep = ?  "
	res, err := con.Db.Queryx(sql, codigo)
	if err != nil {
		log.Fatal("[ERRO] erro ao buscar os dados:", err.Error())
		return
	}

	defer res.Close()
	var cep []*model.DadosCep
	c := new(model.DadosCep)
	for res.Next() {
		err = res.StructScan(&c)
		if err != nil {
			log.Fatal("[ERRO] erro ao buscar os dados:", err.Error())
			return
		}
		cep = append(cep, c)
	}
	cepJSON, _ = json.Marshal(cep)
	return cepJSON
}

//CheckCEP busca os dados selecionados
func CheckCEP(cod string) (id string) {
	fmt.Println(cod)
	cod = "30626670"
	sql := "SELECT id FROM address WHERE cep =" + cod + " LIMIT 1"
	res, err := con.Db.Queryx(sql, cod)
	if err != nil {
		log.Fatal("[ERRO] erro ao buscar os dados:", sql, " - ", err.Error())
	}

	defer res.Close()
	for res.Next() {
		err = res.Scan(&id)
		if err != nil {
			log.Fatal("[ERRO] erro ao buscar os dados:", err.Error())
			return
		}
	}
	return id

}
