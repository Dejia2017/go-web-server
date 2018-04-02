package controllers

import (
	"github.com/astaxie/beego"
)

type WSMainController struct {
	beego.Controller
}

func (c *WSMainController) Get() {
	c.TplName = "ws_index.html"
}