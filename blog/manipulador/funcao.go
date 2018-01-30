package manipulador

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

//Funcao é um manipulador
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando funcao num pacote")
}

// SendMail envia um email
func SendMail(body, assunto string) {
	from := "diego@wsitebrasil.com.br"
	to := "tec.infor321@gmail.com"
	to1 := []string{"tec.infor321@gmail.com"}
	// Configurando as altenticações de e-mail
	auth := smtp.PlainAuth("", "diego@wsitebrasil.com.br", "di@#12SD", "zmail.grupoparanet.com.br")
	// Conectar ao servidor, autenticar, configurar o remetente e o destinatario,
	// e envie o email tudo em uma unica etapa

	dados := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + assunto + " !\r\n" +
		body

	msg := []byte(dados)
	err := smtp.SendMail("zmail.grupoparanet.com.br:587", auth, "diego@wsitebrasil.com.br", to1, msg)
	if err != nil {
		log.Printf("[SENDMAIL] smtp error: %s", err.Error())
		return
	}
}
