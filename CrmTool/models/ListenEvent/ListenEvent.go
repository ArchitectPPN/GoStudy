package ListenEvent

import "github.com/astaxie/beego/orm"

type ListenEvent struct {
	Id              string `orm:"column(id);pk"`
	EventType       string `orm:"column(event_type)"`
	EventKey        string `orm:"column(event_key)"`
	CancelKey       string `orm:"column(cancel_key)"`
	RestoreKey      string `orm:"column(restore_key)"`
	ListenParam     string `orm:"column(listen_param)"`
	ListenStatus    string `orm:"column(listen_status)"`
	DealClassName   string `orm:"column(deal_class_name)"`
	EventParam      string `orm:"column(event_param)"`
	Expire          string `orm:"column(expire)"`
	Synchronization string `orm:"column(synchronization)"`
	Deleted         string `orm:"column(deleted)"`
	DateCreated     string `orm:"column(date_created)"`
	DateModified    string `orm:"column(date_modified)"`
}

func init() {
	orm.RegisterModel(new(ListenEvent))
}
