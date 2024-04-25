package TaskInfo

import "github.com/astaxie/beego/orm"

type EaTasksCstm struct {
	Id        string `orm:"column(id_c);pk"`
	ReceiptId string `orm:"column(ea_receipt_id_c)"`
	StageId   string `orm:"column(ea_stage_id_c)"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(EaTasksCstm))
}
