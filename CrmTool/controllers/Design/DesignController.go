package DesignController

import (
	"CrmTool/Request/DesignRequest"
	"github.com/astaxie/beego"
)

type DesignController struct {
	beego.Controller
}

// CreateDesignForCaseLastStage 为病例最新阶段创建设计
func (c *DesignController) CreateDesignForCaseLastStage() {
	caseCode := c.Ctx.Input.Query("case_code")

	designRequest := DesignRequest.DesignRequest{CaseCode: caseCode}

	c.Ctx.WriteString("设计方案: " + designRequest.CaseCode)
}
