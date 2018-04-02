package controllers

import (
	"fmt"
	//"github.com/astaxie/beego"
	"go-web-server/models"
	"go-web-server/enums"
)

type IndexController struct {
	BaseController
}

func (index *IndexController) Get() {
	index.TplName = "index.tpl"
}

func (index *IndexController) Post() {
	var user models.User
	inputs := index.Input()

	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")

	err := models.Validation(user)

	if err == nil {
		index.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}

func (index *IndexController) Save(){
	index.jsonResult(enums.JRcodeSucc,"success",0)
}