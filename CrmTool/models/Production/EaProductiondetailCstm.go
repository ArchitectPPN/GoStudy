package Production

import "github.com/astaxie/beego/orm"

type EaProductiondetailCstm struct {
	Id                string `orm:"column(id_c);pk"`
	IsBuchaBcq        string `orm:"column(is_bucha_bcq_c)"`
	UpperTimes        string `orm:"column(upper_times_c)"`
	ModelType         string `orm:"column(model_type_c)"`
	Product           string `orm:"column(product_c)"`
	UpperJawEndStep   string `orm:"column(upper_jaw_end_step_c)"`
	Thickness         string `orm:"column(thickness_c)"`
	LowerTimes        string `orm:"column(lower_times_c)"`
	LowerJawEndStep   string `orm:"column(lower_jaw_end_step_c)"`
	UpperJawBeginStep string `orm:"column(upper_jaw_begin_step_c)"`
	LowerJawBeginStep string `orm:"column(lower_jaw_begin_step_c)"`
	Material          string `orm:"column(material_c)"`
	CrudeMaterial     string `orm:"column(crude_material_c)"`
	WearStep          string `orm:"column(wear_step_c)"`
	ProductType       string `orm:"column(product_type_c)"`
	Scheme            string `orm:"column(scheme_c)"`
	Tag               string `orm:"column(tag_c)"`
	Cnt               string `orm:"column(cnt_c)"`
}

func init() {
	orm.RegisterModel(new(EaProductiondetailCstm))
}
