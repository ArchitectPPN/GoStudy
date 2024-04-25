package Stage

import "github.com/astaxie/beego/orm"

type EaStageCstm struct {
	StageId               string `orm:"column(id_c);pk"`
	StageUseA6OrAngelPole string `orm:"column(low_jaw_forward_a6_stage_c)"`
}

func init() {
	orm.RegisterModel(new(EaStageCstm))
}
