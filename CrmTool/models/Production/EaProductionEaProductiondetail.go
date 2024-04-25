package Production

import "github.com/astaxie/beego/orm"

type EaProductionEaProductiondetail_1_c struct {
	Id                  string `orm:"column(id);pk"`
	DateModified        string `orm:"column(date_modified)"`
	Deleted             string `orm:"column(deleted)"`
	ProductionIda       string `orm:"column(ea_production_ea_productiondetail_1ea_production_ida)"`
	ProductionDetailIdb string `orm:"column(ea_production_ea_productiondetail_1ea_productiondetail_idb)"`
}

func init() {
	orm.RegisterModel(new(EaProductionEaProductiondetail_1_c))
}
