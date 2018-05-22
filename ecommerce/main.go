package main

import (
	"fmt"

	"github.com/DiegoSantosWS/ecommerce/autentication"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	t := autentication.GenerateToken()

	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte("QWEFDFSDF"), nil
	})

	if token.Valid == true {
		fmt.Println(t)
	}

	//routers.Routers()
}
