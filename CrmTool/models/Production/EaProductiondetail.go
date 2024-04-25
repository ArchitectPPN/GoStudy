package Production

import "github.com/astaxie/beego/orm"

type EaProductiondetail struct {
	Id             string `orm:"column(id);pk"`
	Name           string `orm:"column(name)"`
	DateEntered    string `orm:"column(date_entered)"`
	DateModified   string `orm:"column(date_modified)"`
	Description    string `orm:"column(description)"`
	ModifiedUserId string `orm:"column(modified_user_id)"`
	CreatedBy      string `orm:"column(created_by)"`
	AssignedUserId string `orm:"column(assigned_user_id)"`
}

func init() {
	orm.RegisterModel(new(EaProductiondetail))
}
