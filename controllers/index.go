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
	//sess := index.StartSession()
	//username := sess.Get("username")
	index.TplName = "index.tpl"
	//fmt.Println(username)
	//if username == nil || username == "" {
	//	index.TplName = "index.tpl"
	//} else {
	//	index.TplName = "success.tpl"
	//}

}

func (index *IndexController) Post() {

	//sess := index.StartSession()
	var user models.User
	inputs := index.Input()


	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")

	err := models.Validation(user)

	//err := models.ValidateUser(user)
	if err == nil {
		//sess.Set("username", user.Username)
		//fmt.Println("username:", sess.Get("username"))
		index.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}

func (index *IndexController) Save(){
	index.jsonResult(enums.JRcodeSucc,"success",0)
}