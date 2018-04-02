package controllers

import (

//"github.com/astaxie/beego"
"go-web-server/enums"
)

type TestController struct {
	BaseController
}

func (index *TestController) Save(){
	//index.TplName = "test.tpl"
	index.jsonResult(enums.JRcodeSucc,"success",0)
}
