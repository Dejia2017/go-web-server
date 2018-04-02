package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

func Jwt_test(){
	mySigningKey := []byte("dejia123") //签名用的秘钥
	// Create the Claims
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //根据claims就可以new一个token 头部？负荷？
	ss, err := token.SignedString(mySigningKey) //用秘钥对 负荷和头部 进行签名 ss 把签名结果也附上
	fmt.Println("签名后的token信息:", ss)

	t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)
}
