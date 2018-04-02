package main

import (
	_ "go-web-server/routers"
	_ "go-web-server/sysinit"
	"github.com/astaxie/beego"
	//"go-web-server/utils"
	//"fmt"
	//"go-web-server/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

func main() {
	mySigningKey := []byte("hzwy23") //字符数组
	// Create the Claims
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //根据claims就可以new一个token
	ss, err := token.SignedString(mySigningKey) //
	fmt.Println("签名后的token信息:", ss)
	t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)

	//utils.UtilTest()

	beego.Run()
}


