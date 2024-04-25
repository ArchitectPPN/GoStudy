package routers

import (
	TaskController "CrmTool/controllers/Task"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/getTaskExternal", &TaskController.TaskController{}, "get:GetTaskExternalByTaskCode")
	beego.Router("/getTaskReceipt", &TaskController.TaskController{}, "get:GetTaskReceiptByTaskCode")
	beego.Router("/GetTaskMesQcMsg", &TaskController.TaskController{}, "get:GetTaskMesQcResultJsonByTaskCode")
	beego.Router("/GetSilicaQcResult", &TaskController.TaskController{}, "get:GetSilicaQcResult")
	beego.Router("/testDemo", &TaskController.TaskController{}, "Get:TestDemo")

	// 上传stl地址
	beego.Router("/GetUploadStlResult", &TaskController.TaskController{}, "get:GetUploadStlResult")
}
