package ProductionController

import (
	"CrmTool/Request/ProductionRequest"
	"CrmTool/Services/ProductionService"
	"github.com/astaxie/beego"
)

type ProductionController struct {
	beego.Controller
}

func (p *ProductionController) CreateProductionForStage() {
	var productionRequest ProductionRequest.ProductionRequest
	productionRequest = ProductionRequest.ProductionRequest{}

	// 初始化阶段id
	productionRequest.StageId = p.Ctx.Input.Query("stage_id")
	// 初始化产品种类
	productionRequest.ApplianceType = p.Ctx.Input.Query("appliance")

	createProductionService := new(ProductionService.CreateProductionService)
	createProductionService.CreateProduction(&productionRequest)

	p.Ctx.WriteString("输入的阶段id为: " + productionRequest.StageId)
}
