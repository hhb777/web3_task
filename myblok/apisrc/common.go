package apisrc

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("hhbuser")

func TokenAuth(stoken string) (res bool, err error) {
	restoken, err := jwt.Parse(stoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名方法不正确")
		}
		return JwtKey, nil
	})
	if err != nil || !restoken.Valid {
		return false, err
	}
	return true, nil
}
