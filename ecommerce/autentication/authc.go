package autentication

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken() string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "lksdjajduieyiuHGFGFDSFSER",
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenstring, err := token.SignedString([]byte("QWEFDFSDF"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func ValidateTokenMidllers() {

}
