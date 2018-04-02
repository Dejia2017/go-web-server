package sysinit

import (
	"github.com/astaxie/beego"
	"go-web-server/utils"
)

//初始化数据连接
func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	utils.InitLogs()
	//初始化缓存
	utils.InitCache()
	//初始化数据库
	InitDatabase()
}
