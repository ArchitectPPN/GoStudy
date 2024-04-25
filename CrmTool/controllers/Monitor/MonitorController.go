package MonitorController

import (
	"CrmTool/Services/ListenEventService"
	"github.com/astaxie/beego"
)

type MonitorController struct {
	beego.Controller
}

func (c *MonitorController) GetWaitModelStageEventList() {
	ListenEventServiceObj := new(ListenEventService.ListenEventService)

	list := ListenEventServiceObj.GetWaitEventList()
	var eventList, eventType string
	eventList += "<table><tr><td>类型</td><td>等待触发天数</td></tr>"
	for _, val := range list {
		switch val[0].(string) {
		case "wait_new_stage_for_production":
			eventType = "等待阶段"
		default:
			eventType = "等待模型"
		}

		eventList += "<tr><td>" + eventType + "</td><td>" + val[3].(string) + " 天</td></tr>"
	}

	eventList += "</table>"
	c.Ctx.WriteString(eventList)
}
