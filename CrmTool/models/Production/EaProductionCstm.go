package Production

import "github.com/astaxie/beego/orm"

type EaProductionCstm struct {
	Id                  string `orm:"column(id_c);pk"`
	State               string `orm:"column(state_c)"`
	ExpireDate          string `orm:"column(expire_date_c)"`
	EaCaseId            string `orm:"column(ea_case_id_c)"`
	Products            string `orm:"column(products_c)"`
	IsFollow            string `orm:"column(is_follow_c)"`
	OdsAddr             string `orm:"column(ods_addr_c)"`
	Num                 string `orm:"column(num_c)"`
	ProduceRemark       string `orm:"column(produce_remark_c)"`
	Province            string `orm:"column(province_c)"`
	City                string `orm:"column(city_c)"`
	Area                string `orm:"column(area_c)"`
	PostalCode          string `orm:"column(postalcode_c)"`
	Contact             string `orm:"column(contact_c)"`
	Telephone           string `orm:"column(telephone_c)"`
	Address             string `orm:"column(address_c)"`
	ModelType           string `orm:"column(model_type_c)"`
	CaseType            string `orm:"column(case_type_c)"`
	Tag                 string `orm:"column(tag_c)"`
	CFrom               string `orm:"column(c_from_c)"`
	Jaw                 string `orm:"column(jaw_c)"`
	Proportion          string `orm:"column(proportion_c)"`
	UpperSiliconBarcode string `orm:"column(upper_silicon_barcode_c)"`
	LowerSiliconBarcode string `orm:"column(lower_silicon_barcode_c)"`
}

func init() {
	orm.RegisterModel(new(EaProductionCstm))
}
