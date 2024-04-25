package Case

import "github.com/astaxie/beego/orm"

type EaCase struct {
	Id                string `orm:"column(id);pk"`
	Name              string `orm:"column(name)"`
	DateEntered       string `orm:"column(date_entered)"`
	DateModified      string `orm:"column(date_modified)"`
	ModifiedUserId    string `orm:"column(modified_user_id)"`
	CreatedBy         string `orm:"column(created_by)"`
	Description       string `orm:"column(description)"`
	Deleted           string `orm:"column(deleted)"`
	AssignedUserId    string `orm:"column(assigned_user_id)"`
	RegionId          string `orm:"column(region_id)"`
	OverdueDate       string `orm:"column(overdue_date)"`
	ProductionRuleTag string `orm:"column(production_rule_tag)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(EaCase))
}
