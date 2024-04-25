package QcResult

type QcDetail struct {
	// 模型
	CouldBuccinator int         `json:"couldBuccinator,omitempty"` // 唇颊肌
	QcObjects       []QcItem    `json:"qcObjects,omitempty"`       // 质检结果详情
	RiskDetail      *RiskDetail `json:"riskDetail,omitempty"`      // 风险详情

	// 不合格url
	MesUnqualifiedReportUrl string `json:"mesUnqualifiedReportUrl,omitempty"`
	// 不合格原因
	UnqualifiedReason string `json:"unqualifiedReason,omitempty"`
	// 不合格code
	UnqualifiedReasonCode string `json:"unqualifiedReasonCode,omitempty"`
	// 缺陷情况描述
	DefectiveRemark string `json:"defectiveRemark,omitempty"`
	// 缺陷部位照片的MES下载地址
	DefectivePartPhotos []string `json:"defectivePartPhotos,omitempty"`

	// 照片
	PhotoDetail *PhotoDetail `json:"photoDetail,omitempty"` // 照片质检结果

	// 公用
	QcItem string `json:"qcItem"`
	QcType string `json:"qcType"` // 质检类型。实物质检合格不合格都要告知。MATERIAL_QC实物质检，DIGITAL_QC数据质检
}

// QcItem 质检结果详情
type QcItem struct {
	BarCode    string `json:"barCode,omitempty"`
	CdsUrl     string `json:"cdsUrl,omitempty"`
	Jaw        string `json:"jaw,omitempty"`
	QcResult   int    `json:"qcResult,omitempty"`
	QcStandard string `json:"qcStandard,omitempty"` // 检测标准（notA6:非A6病例检测标准，A6:A6病例检测标准） 当barcode为咬合记录条形码，且A6病例时，本字段必然是A6，其余情况都为notA6
}

// RiskDetail 风险详情
type RiskDetail struct {
	PromptRisk     int      `json:"promptRisk"`
	RiskDesc       string   `json:"riskDesc,omitempty"`
	RiskDescPhotos []string `json:"riskDescPhotos,omitempty"`
	RiskReportUrl  string   `json:"riskReportUrl,omitempty"`
}
