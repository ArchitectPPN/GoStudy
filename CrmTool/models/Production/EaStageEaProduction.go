package Production

import "github.com/astaxie/beego/orm"

type EaStageEaProduction_1_c struct {
	Id           string `orm:"column(id);pk"`
	DateModified string `orm:"column(date_modified)"`
	StageId      string `orm:"column(ea_stage_ea_production_1ea_stage_ida)"`
	ProductionId string `orm:"column(ea_stage_ea_production_1ea_production_idb)"`
}

func init() {
	orm.RegisterModel(new(EaStageEaProduction_1_c))
}
