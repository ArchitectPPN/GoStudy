package Case

import "github.com/astaxie/beego/orm"

type EaCaseCstm struct {
	Id            string `orm:"column(id_c);pk"`
	ApplianceType string `orm:"column(appliance_type_c)"`
}

func init() {
	orm.RegisterModel(new(EaCaseCstm))
}
