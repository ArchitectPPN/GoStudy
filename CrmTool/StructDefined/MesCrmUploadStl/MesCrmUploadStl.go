package MesCrmUploadStl

// MesCrmUploadStl 上传stl地址
type MesCrmUploadStl struct {
	UploadResult []Detail `json:"uploadResult"`
	WorkOrderId  string   `json:"workOrderId"`
}

// Detail 详情
type Detail struct {
	CdsUrl             string `json:"cdsUrl,omitempty"`
	GenerateFromRemade bool   `json:"generateFromRemade"`
	Jaw                string `json:"jaw,omitempty"`
	StlSource          string `json:"stlSource,omitempty"`
	BarCode            string `json:"barCode,omitempty"`
}
