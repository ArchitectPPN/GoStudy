package Stage

import "github.com/astaxie/beego/orm"

type EaStage struct {
	Id             string `orm:"column(id);pk"`
	Name           string `orm:"column(name)"`
	DateEntered    string `orm:"column(date_entered)"`
	DateModified   string `orm:"column(date_modified)"`
	ModifiedUserId string `orm:"column(modified_user_id)"`
	CreatedBy      string `orm:"column(created_by)"`
	Description    string `orm:"column(description)"`
	Deleted        int8   `orm:"column(deleted)"`
	AssignedUserUd string `orm:"column(assigned_user_id)"`
}

func init() {
	orm.RegisterModel(new(EaStage))
}
