package routers

import (
	DesignController "CrmTool/controllers/Design"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("createDesignForCase", &DesignController.DesignController{}, "get:CreateDesignForCaseLastStage")
}
