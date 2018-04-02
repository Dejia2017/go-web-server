package controllers

import (
	"fmt"
	//"net/url"
	"go-web-server/models"
)

type RegistController struct {
	BaseController
}

func (this *RegistController) Get() {
	this.TplName = "regist.tpl"
}

func (this *RegistController) Post() {

	var user models.User
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username") //从name中获取
	user.Pwd = inputs.Get("pwd")
	err := models.SaveUser(user)

	if err == nil {
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "error.tpl"
	}
}
