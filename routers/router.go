package routers

import (
	"go-web-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/regist",&controllers.RegistController{})
	beego.Router("/",&controllers.IndexController{})
	beego.Router("/test",&controllers.TestController{},"*:Save")

	beego.Router("/ws",&controllers.MyWebSocketController{})
	beego.Router("/w",&controllers.WSMainController{})

}
