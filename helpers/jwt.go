package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id, secret string)(string,error) {
	mySigningKey := []byte(secret);

	expire:=time.Now().Add(60*time.Minute).Unix()

	claims:= LoginClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expire,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString,err
}


func ValidateToken(secret,tokenString string)(*LoginClaims,error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	

		if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
			return claims,nil
		} else {
			return claims, err
		}
}