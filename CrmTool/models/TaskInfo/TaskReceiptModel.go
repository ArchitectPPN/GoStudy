package TaskInfo

import (
	"github.com/astaxie/beego/orm"
)

type TaskReceipt struct {
	Id             int    `orm:"column(id);pk"`
	TaskId         string `orm:"column(task_id)"`
	ReceiptId      string `orm:"column(receipt_id)"`
	Jaw            string `orm:"column(jaw)"`
	ReceiptType    string `orm:"column(receipt_type)"`
	ReceiptInfo    string `orm:"column(receipt_info)"`
	ReceiptOdsStep int    `orm:"column(receipt_ods_step);default(0);type(int)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(TaskReceipt))
}
