package Stage

import "github.com/astaxie/beego/orm"

type EaCaseEaStage_1_c struct {
	Id           string `orm:"column(id);pk"`
	DateModified string `orm:"column(date_modified)"`
	Deleted      int8   `orm:"column(deleted)"`
	CaseId       string `orm:"column(ea_case_ea_stage_1ea_case_ida)"`
	StageId      string `orm:"column(ea_case_ea_stage_1ea_stage_idb)"`
}

func init() {
	orm.RegisterModel(new(EaCaseEaStage_1_c))
}
