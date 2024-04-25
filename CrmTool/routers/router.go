package routers

import (
	"CrmTool/controllers/Test"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &Test.MainController{})
	beego.Router("/testKafka", &Test.MainController{}, "get:TestConnectKafka")
}
