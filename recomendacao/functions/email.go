package functions

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/smtp"

	con "github.com/DiegoSantosWS/recomendacao/connection"
)

//DadosProdutos os dados que serão enviados para template
type DadosProdutos struct {
	IDP  int            `json:"id" db:"id"`
	Name string         `json:"nome" db:"nome"`
	Img  sql.NullString `json:"img" db:"img"`
}

//DadosEmail struct retorna as informações de emails
type DadosEmail struct {
	Link          string
	Nome          string
	Corpo         string
	Recomendacoes string
	URL           string
	Msg           string
	Att           string
	UltimaCompra  int
}

/*EnviaEmail envia o email para cliente destinado...
 * nomeCliente  string, recebe nome do cliente
 * emailCliente string, recebe email do cliente
 * nomeProduto  string, recebe nome do produto
 * ultimaCompra int,    recebe valor da ultima compra
 * categoria 	int,    recebe o codigo da categoria
 * produto 		int     recebe codigo do produto
 * @return true
 */
func EnviaEmail(nomeCliente, emailCliente, nomeProduto string, ultimaCompra, categoria, produto int) {
	url := "https://www.shoppingdoparafuso.com.br"

	sql := "SELECT p.id, p.nome, "
	sql += "(SELECT g.caminho FROM sig_ged as g WHERE g.modulo = 'Produtos 1.0' AND g.tipo IN('image/jpeg','image/png') AND g.idModulo = p.id ORDER BY g.id DESC LIMIT 1) as img"
	sql += " FROM sig_produtos as p WHERE p.categoria = ? AND p.id != ? AND p.status = 'Liberado' LIMIT 10"
	row, err := con.Db.Queryx(sql, categoria, produto)
	if err != nil {
		fmt.Println("ERROR: Produto não encontrado", sql, err.Error())
		return
	}

	//var Dp DadosProdutos
	//var ps []DadosProdutos
	var local DadosProdutos
	var ps []DadosProdutos
	defer row.Close()
	for row.Next() {
		err = row.StructScan(&local)
		if err != nil {
			fmt.Println("ERROR: Produto não encontrado não encontrado", sql, err.Error())
			return
		}
		ps = append(ps, local)
	}
	//name, _ := NameSlugCategoria(categoria)
	data := map[string]interface{}{
		"produtos":      ps,
		"Nome":          nomeCliente,
		"UltimaCompra":  ultimaCompra,
		"Corpo":         nomeProduto,
		"Recomendacoes": "Então, preparamos algumas opçoes de produtos para você e achamos que irá gostar, vejá a lista abaixo: ",
		"URL":           url,
		"Att":           "https://www.shoppingdoparafuso.com.br/imgs/logo.jpg",
		"Categoria":     "teste-casa-de-casa",
		"Link":          "http://www.casadecasavirtual.com.br",
	}
	r := NewRequest([]string{"viniciustda13@hotmail.com"}, nomeCliente+" - Recomendamos", "Recomendandos", "CASA DE CASA")
	err = r.ParseTemplate("template/sendEmail.html", data)
	if err == nil {
		ok, err := r.EnviarEmail()
		if err != nil {
			fmt.Println("ERROR: não foi possivel enviar o email: ", err.Error())
		}
		fmt.Println(ok)
	} else {
		fmt.Println("Diego santos", err.Error())
	}

}

//Request requerendo os dados para struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

/*NewRequest preenche a struct com dados do cabeçalho do email e corpo da mensagem...
 * to      []string recebe um slice com emails do cliente,
 * subject string recebe o valor do assunto do email,
 * body    string recebe o corpo do email,
 * from    string valor do from
 * @return retorna um ponteiro de Request
 */
func NewRequest(to []string, subject, body, from string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		from:    from,
		body:    body,
	}
}

/*EnviarEmail recebe um ponteiro da requisição NewRequest
 * para enviar email destinado ao cliente
 */
func (r *Request) EnviarEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := "FROM: " + r.from + "!\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(from + subject + mime + "\n" + r.body)

	auth := smtp.PlainAuth("", "diego@wsitebrasil.com.br", "di@#12SD", "zmail.grupoparanet.com.br")
	err := smtp.SendMail("zmail.grupoparanet.com.br:587", auth, "diego@wsitebrasil.com.br", r.to, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}

//ParseTemplate executa a template que vai compor o corpo de email.
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err.Error())
		return err
	}
	r.body = buf.String()
	return nil
}
