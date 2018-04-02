package main

import (
	_ "go-web-server/routers"
	_ "go-web-server/sysinit"
	"github.com/astaxie/beego"
	//"go-web-server/utils"
	"go-web-server/utils"
)

func main() {


	//utils.UtilTest()
	utils.Jwt_test()

	beego.Run()
}


