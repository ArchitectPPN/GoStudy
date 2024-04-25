package TaskInfo

import "github.com/astaxie/beego/orm"

type EaTasks struct {
	Id             string `orm:"column(id);pk"`
	Name           string `orm:"column(name)"`
	Deleted        int    `orm:"column(deleted)"`
	FileName       string `orm:"column(filename)"`
	CreatedBy      string `orm:"column(created_by)"`
	Description    string `orm:"column(description)"`
	FileMimeType   string `orm:"column(file_mime_type)"`
	AssignedUserId string `orm:"column(assigned_user_id)"`
	ModifiedUserId string `orm:"column(modified_user_id)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(EaTasks))
}
