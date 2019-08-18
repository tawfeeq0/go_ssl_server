package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tawfeeq0/go_ssl_server/server/config"
)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRETKEY)

}
