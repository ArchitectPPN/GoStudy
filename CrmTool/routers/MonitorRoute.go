package routers

import (
	MonitorController "CrmTool/controllers/Monitor"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("getWaitModelStageEventList", &MonitorController.MonitorController{}, "get:GetWaitModelStageEventList")
}
