package controllers

import (
	"github.com/astaxie/beego"
	"go-web-server/enums"
	"go-web-server/models"
)

type BaseController struct {
	beego.Controller
}

//使用json作为数据交互的格式
func (c *BaseController) jsonResult(code enums.JsonResultCode,msg string, obj interface{}) {
	r := &models.JsonResult{code,msg,obj}
	c.Data["json"] = r
	c.ServeJSON() //ServeJSON sends a json response with encoding charset.
	c.StopRun()
}