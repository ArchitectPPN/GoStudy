package ListenEventService

import (
	"github.com/astaxie/beego/orm"
)

type ListenEventService struct {
}

// ListenEventList ListenEventList
type ListenEventList struct {
	Id           string
	EventKey     string
	ListenStatus int8
	DateCreated  string
	DateDiff     string
}

// GetWaitEventList GetWaitEventList
func (service *ListenEventService) GetWaitEventList() []orm.ParamsList {
	var list []orm.ParamsList

	ormObj := orm.NewOrm()
	_, _ = ormObj.Raw("SELECT event_type, listen_status, CONVERT_TZ(date_created, 'UTC', 'Asia/Shanghai') as date_created, DATEDIFF(NOW(), date_created) as date_diff FROM listen_event WHERE listen_status = ? AND event_type IN (?, ?)", 0, "wait_model_for_production", "wait_new_stage_for_production").ValuesList(&list)

	return list
}
