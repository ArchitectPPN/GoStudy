package routers

import (
	ProductionController "CrmTool/controllers/Production"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("createProductionForStage", &ProductionController.ProductionController{}, "get:CreateProductionForStage")
}
